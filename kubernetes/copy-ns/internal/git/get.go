package git

import (
  "fmt"
  "os"
  "strings"
  "encoding/base64"
)



func getProjectID(projectName string) (string, error) {
  projects := httpGet(os.Getenv("GITLAB_URL")+"/api/v4/groups/"+os.Getenv("GITLAB_GROUP_ID")+"/search?scope=projects&search="+projectName)
  if len(projects) != 1 {
    for _, p := range projects {
      if strings.HasSuffix(p.Get("name"), projectName){
        return p.Get("id"), nil
      }
    }
    err := fmt.Errorf("Multiples projects founded for "+projectName)
    return "", err
  }
  return projects[0].Get("id"), nil
}

func GetFileSha(projectID string, fileName string, defaultFileName string, path string) (string) {
  sha := ""
  files := httpGet(os.Getenv("GITLAB_URL")+"/api/v4/projects/"+projectID+"/repository/tree?path="+path)
  for _, f := range files {
    if(fileName == f.Get("name")){
      return f.Get("id")
    } else if(defaultFileName == f.Get("name")){
      sha = f.Get("id")
    }
  }
  if sha == "" {
    panic("File '"+fileName+"' or '"+defaultFileName+"' not found for "+projectID)
  }
  return sha
}

func GetFile(prjectID string, fileSha string) (string) {
  blob := httpGet(os.Getenv("GITLAB_URL")+"/api/v4/projects/"+prjectID+"/repository/blobs/"+fileSha)
  file, err := base64.StdEncoding.DecodeString(blob[0].Get("content"))
  if err != nil {
      panic(err)
  }
  return string(file)
}

func GetEnvFile(projectName string, environment string) (string, error){
  projectID, err := getProjectID(projectName)
  if err != nil {
    return "", err
  }
  fileSha := GetFileSha(projectID, environment+".yaml", "config.yaml", "env")
  env := GetFile(projectID, fileSha)
  return env, nil
}