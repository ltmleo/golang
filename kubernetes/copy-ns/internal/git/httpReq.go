package git

import (
	"encoding/json"
	"strings"
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
  )

type response map[string]interface{}
type responseList []response

func (r response) Get(str string) string {
	return fmt.Sprintf("%v", r[str])
}

func httpGet(url string) (responseList) {
	var data responseList
	method := "GET"
  
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(``))
  
	if err != nil {
	  panic(err)
	}
	req.Header.Add("PRIVATE-TOKEN", os.Getenv("GITLAB_API_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
  
	res, err := client.Do(req)
	if err != nil {
	  panic(err)
	}
	defer res.Body.Close()
  
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	  panic(err)
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		var dataAux response
		err = json.Unmarshal(body, &dataAux)
		if err != nil {
			panic(err)
		}
		data = responseList{dataAux}
	}
	return data
}