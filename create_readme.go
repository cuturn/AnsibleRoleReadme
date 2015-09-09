//[current directory]/roles/...
package main

import (
   "gopkg.in/yaml.v2"
   "os"
   "io/ioutil"
   "github.com/kr/pretty"
    "fmt"
    "sort"
)

type FileInfos []os.FileInfo
type ByName struct{ FileInfos }
func (fi ByName) Len() int {
        return len(fi.FileInfos)
}
func (fi ByName) Swap(i, j int) {
        fi.FileInfos[i], fi.FileInfos[j] = fi.FileInfos[j], fi.FileInfos[i]
}
func (fi ByName) Less(i, j int) bool {
        return fi.FileInfos[j].ModTime().Unix() < fi.FileInfos[i].ModTime().Unix()
}

func isExist(filename string) bool {
    _, err := os.Stat(filename)
        return err == nil
}

func main() {
  var roleDir, _ = os.Getwd()
  roleDir += "/roles/"
  fileInfos, err := ioutil.ReadDir(roleDir)

  if err != nil {
    fmt.Errorf("Directory cannot read %s\n", err)
    os.Exit(1)
  }
  
  var readme = ""
  var readme_filepath = ""
  var defaults_filepath = ""
  sort.Sort(ByName{fileInfos})
  for _ , fileInfo := range fileInfos {
    if fileInfo.IsDir() {
       readme = "#" + fileInfo.Name() + "\n\n"
       readme += "## abstruct \n\n"
       readme += "## variables \n\n"
       defaults_filepath = roleDir + fileInfo.Name() + "/defaults/main.yml"
       
       source,err := ioutil.ReadFile(defaults_filepath)
       if err != nil {
         fmt.Printf("%s\n", defaults_filepath + ":file cannot open.")
         continue
       }
       m := make(map[interface{}]interface{})
       err = yaml.Unmarshal(source, &m)
       if err != nil {
         fmt.Printf("%s\n", defaults_filepath + ":fail to parse yaml file.")
	 continue
       }
       
       readme += "|variable name|default value|\n|---|---|\n"
       for index,element := range m{
          readme += "|" + fmt.Sprint(index) + "|" + fmt.Sprint(element) + "| \n"
       }
       readme_filepath = roleDir + fileInfo.Name() + "/README.md"
       err = ioutil.WriteFile(readme_filepath,[]byte(readme),0644)
       if err != nil {
         fmt.Printf("%s\n" , readme_filepath + ":fail to write readme.")
       }
       pretty.Printf("--- m:\n%# v\n\n", m)
    }
  }
}

