package json

import (
	"encoding/json"
	"testing"
)

var jsonStr = `{
  "basic_info":{
  	"name":"wuzhixuan",
      "age":24
  },
    "job_info":{
    	"skills":["java","go","C"]
    }
}`

func TestEmbeddedJson(t *testing.T) {

	e := new(Employee)

	json.Unmarshal([]byte(jsonStr), e)
	t.Log(e)

	str, err := json.Marshal(e)

	t.Log(string(str), err)

}
