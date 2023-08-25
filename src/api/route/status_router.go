package route

import (
	"github.com/josewenzel/flagify/src/api/handler"
	"net/http"
)

func StatusRequest() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			handler.HandleGetStatus(writer, request)
		default:
			handler.StatusRespondWith(http.StatusNotImplemented, writer)
		}
	}
}
