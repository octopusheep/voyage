package main

import(
	"io"
	"net/http"
	"github.com/octopusheep/voyage/api/defs"

)

func sendErrorResponse(w http.ResponseWriter,errResp defs.ErrResponse){
	w.WriteHeader(errResp.HttpSC)

	resStr,_:=json.Marshal(&errResp.Error)
	io.WriterString(w,string(resStr))
}

func sendNormalResponse(w http.ResponseWriter,resp string,sc int){
	w.WriteHeader(sc)
	io.WriteString(w,resp)
}
