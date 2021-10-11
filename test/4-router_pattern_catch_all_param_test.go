package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouterPatternCatchAllParameters(t *testing.T) {
	router := httprouter.New()

	//with anonymous functiom
	router.GET("/images/*image", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		image := p.ByName("image")
		text := "Image: " + image

		fmt.Fprint(rw, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/images/small/sad/oli.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image: /small/sad/oli.png", string(body))
}
