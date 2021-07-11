package wfw

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

type Person struct {
	Name string `validate:"required"`
	Address string `validate:"required"`
	Age int `validate:"required"`
}
var personResponse = Person {
		Name : "tom",
		Address : "shanghai",
		Age : 20,
	}


var personResponseBytes, _ = json.Marshal(personResponse)

func TestJsonBinding_Bind(t *testing.T) {
	req := httptest.NewRequest("POST", "/test", bytes.NewBuffer(personResponseBytes))
	req.Header.Set("Content-Type","application/json")
	var p Person
	j :=jsonBinding{}
	err := j.Bind(req,&p)
	if err != nil {
		t.Fatal("fail",err.Error())
	}
	if !assert.Equal(t,p,personResponse){
		t.Fatal("fail")
	}


}
