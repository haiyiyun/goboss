package service1

import (
	"net/http"
)

func (self *Service) Route_GET_Service1(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("service1"))

}
