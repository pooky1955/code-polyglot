package main

import (
    "bufio"
    "flag"
    "fmt"
    _ "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "sort"
    "strings"
    _ "strings"
)

//extension_map is a map of file extensions to their respective language names
var extension_map map[string]string = map[string]string{
    "hs"  : "Haskell",
    "go" :  "Go",
    "py" : "Python",
}

//check will invoke log.Fatal if err != nil 
func check(err error){
    if err != nil {
        log.Fatal(err)
    }
}


//Filepath is a struct that stores the direcotry, the filename, and the extension.
type Filepath struct {
    directory string
    filename string
    extension string
}

//splitFileName takes a full_filename like foo.hs and returns ("foo","hs",nil)
//in the case of an error like foo, it returns ("","",error "Could not split filename foo")
func splitFileName(full_filename string) (string,string,error){
    splitted := strings.Split(full_filename,".")
    if len_splitted := len(splitted); len_splitted > 1 {
        filename := strings.Join(splitted[0:len_splitted-1],".")
        extension := splitted[len_splitted-1]
        return filename,extension, nil
    } else {
        return "", "", fmt.Errorf("Could not split filename %v",full_filename)
    }
}

//getAllFiles gets all files in the current directory and returns their filenames 
//relative to the directory passed as parameter
func getAllFiles(recursive bool,directory string) []Filepath {
    // read all files and folders
    files, err := ioutil.ReadDir(directory)
    check(err)
    all_files := make([]Filepath,0,len(files))

    for _, file := range files {
        full_path := filepath.Join(directory,file.Name())
        if file.IsDir() {
            // if file is directory, recursively traverse it
            if recursive {
                all_files = append(all_files,getAllFiles(true,full_path)...)
            }
        } else {
            full_filename := file.Name()
            filename,extension,err := splitFileName(full_filename)
            if err == nil {
                filepath := Filepath{directory : directory, filename : filename, extension : extension}
                all_files = append(all_files,filepath)
            }
        }
    }
    fmt.Println(all_files)
    return all_files

}


//group_files groups paths that have the same filepath except for extension under the same key
//so given as input ['foo/bar/baz.hs','foo/bar/baz.go','foo/bar/boo.py']
//it outputs {'foo/bar/baz' : ['hs','go'],'foo/bar/boo' : 'py'}
func group_files(paths []Filepath) map[string][]string{
    groups := make(map[string][]string)
    for _, path := range paths {
        dir, fn, ext := path.directory, path.filename, path.extension
        full_path := filepath.Join(dir,fn)
        fmt.Println("full path from groupfiles",full_path)
        fmt.Println(path)
        if ext_slice, ok := groups[full_path]; ok {
            groups[full_path] = append(ext_slice,ext)
        } else {
            groups[full_path] = []string{ext}
        }
    }
    return groups
}

//generate_toc generates a table of contents in markdown in the form of 
// 1. link1
// 2. link2  
// 3. link3
func generate_toc(exts []string,extension_dir map[string]string) string{
    languages := make([]string,len(exts))
    for i,ext := range exts {
        languages[i] = extension_dir[ext]
    }
    sb  := strings.Builder{}
    sb.WriteString("# Table of contents\n")
    for i, lang := range languages {
        lang_link := fmt.Sprintf("%d. [%v](#%v)\n",i,lang,lang)
        sb.WriteString(lang_link)
    }
    result := sb.String()
    return result
}

//capitalize takes a string foo and capitalizes it to Foo
func capitalize(word string) string {
    sb := strings.Builder{}
    sb.WriteString(strings.ToUpper(string(word[0])))
    sb.WriteString(word[1:])
    return sb.String()

}
//pretty_parse converts a filename foo-bar-baz to Foo bar baz
func pretty_parse(filename string) string{
    splitted := strings.Split(filename,"-")
    sb := strings.Builder{}
    if len_splitted := len(splitted); len_splitted > 1 {
        sb.WriteString(capitalize(splitted[0]))
        sb.WriteString(" ")
        sb.WriteString(strings.Join(splitted[1:]," "))
    }
    return sb.String()

}

//generate_title takes a path as input like foo/bar-baz-boo and returns # Bar baz boo
func generate_title(full_path string) string {
    _, filename := filepath.Split(full_path)
    title := pretty_parse(filename)

    result := fmt.Sprintf("# %v\n",title)
    return result
}

//generate_markdown takes the groups of files and the output directory
//and converts each group to one markdown file with a structure like
//# Title here
//## Toc
//1. link1
//2. link2
//3. etc.
//### code language here  
//```code here```
//etc
func generate_markdown(groups map[string][]string,output_dir string){

    for full_path, exts := range groups {
        //build markdown content
        md_sb := strings.Builder{}
        sort.Strings(exts)

        toc_md := generate_toc(exts,extension_map)
        title_md := generate_title(full_path)

        md_sb.WriteString(title_md)
        md_sb.WriteString(toc_md)

        for _, ext := range exts {
            if ext_lang,ok := extension_map[ext]; ok {
                complete_path := fmt.Sprintf("%v.%v",full_path,ext)
                content,err := ioutil.ReadFile(complete_path)

                if err == nil {
                    top_markdown := fmt.Sprintf("\n## %v\n```%v\n",ext_lang,ext)
                    bottom_markdown := "\n```\n"
                    md_sb.WriteString(top_markdown)
                    md_sb.Write(content)
                    md_sb.WriteString(bottom_markdown)
                }

            }
        }
        markdown := []byte(md_sb.String())
        fmt.Println("full_path=",full_path)
        markdown_file := filepath.Join(output_dir,full_path+".md")
        markdown_dir,_ := filepath.Split(markdown_file)
        os.MkdirAll(markdown_dir,os.ModePerm)
        err := ioutil.WriteFile(markdown_file,markdown,0644) // -rw-rw-r--
        if err != nil {
            log.Fatalf("Error with generating %v: %v",markdown_file,err)
        }
    }

}

func main(){
    //CLI SETUP
    var directory string
    var recursive bool
    var output_dir string

    flag.StringVar(&directory,"directory",".","specifies directory to generate markdown from. defaults to current working directory")
    flag.StringVar(&output_dir,"output","./markdown/","specifies output directory for markdown. defaults to ./markdown/")
    flag.BoolVar(&recursive,"recursive",true,"specifies wheter it should generate markdown from all subdirectories recursively. defaults to true")
    flag.Parse()

    full_path,err := filepath.Abs(directory)
    check(err)

    //CONFIRMING CLI PARAMS
    fmt.Printf("Generating markdown files from %v to %v with recursive=%t\n",full_path,output_dir,recursive)
    fmt.Printf("Do you want to proceed? y/[n]")
    reader := bufio.NewReader(os.Stdin)
    if str,err := reader.ReadString('\n'); err != nil || strings.ToLower(str)[0] != 'y' {
        fmt.Printf("str was equal to %v\n",str)
        fmt.Printf("err was equal to %v\n",err)
        log.Fatal("Aborting program")
    }


    //MAIN CODE
    all_files := getAllFiles(recursive,directory)
    grouped_files := group_files(all_files)
    generate_markdown(grouped_files,output_dir)
}
