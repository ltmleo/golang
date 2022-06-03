package git

import (
	"encoding/json"
	"fmt"
  "os"
	"io/ioutil"
	"net/http"
	"bytes"
)

type payload struct {
	Branch string `json:"branch"`
	CommitMessage string `json:"commit_message"`
	Actions []Action `json:"actions"`
}

type Action struct {
	Action string `json:"action"`
	FilePath string `json:"file_path"`
	Content string `json:"content"`
}

func (p payload) ToBuffer() *bytes.Buffer {
	pJson, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(pJson)
}

func Commit(actions []Action) string {

  url := os.Getenv("GITLAB_URL")+"/api/v4/projects/"+os.Getenv("GITLAB_GROUP_ID")+"/repository/commits/"
  method := "POST"

  p := payload{Branch: "master", CommitMessage: os.Getenv("GIT_COMMIT_MESSAGE"), Actions: actions }

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, p.ToBuffer())

  if err != nil {
    fmt.Println(err)
    return "Commit Error"
  }
  req.Header.Add("PRIVATE-TOKEN", os.Getenv("GITLAB_API_TOKEN"))
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return "Commit Error"
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return "Commit Error"
  }
  return string(body)
}