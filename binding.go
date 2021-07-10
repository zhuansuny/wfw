package wfw

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)
// Content-Type MIME of the most common data formats.
const (
	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
	MIMEPROTOBUF          = "application/x-protobuf"
	MIMEMSGPACK           = "application/x-msgpack"
	MIMEMSGPACK2          = "application/msgpack"
	MIMEYAML              = "application/x-yaml"
)
var (
	JSON          = jsonBinding{}

)
// 仅实现POST JOSN方式
type Binding interface {
	Name()	string
	Bind(*http.Request,interface{})error
}



type jsonBinding struct{}

func (j jsonBinding)Name()string{
	return "json"
}

func (j jsonBinding)Bind(req *http.Request,obj interface{})error{
	if req == nil || req.Body == nil {
		return errors.New("invalid request")
	}
	return  decodeJSON(req.Body, obj)
}

func decodeJSON(r io.Reader,obj interface{})error{
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return  Validator(obj)  //校验
}

func Default(method, contentType string) Binding {
	switch contentType {
	case MIMEJSON:
		return JSON
	//case MIMEXML, MIMEXML2:
	//	return XML
	default: // case MIMEPOSTForm:
		return JSON
	}
}