package testing

import (
	"net/http"
	// "net/http/httptest"
	"reflect"
	"testing"

	// "github.com/gin-gonic/gin"
	// customerControllers "github.com/putragabrielll/fwg17-cinematix-be/src/controllers/customer"
	"github.com/putragabrielll/fwg17-cinematix-be/src/lib"
	"github.com/putragabrielll/fwg17-cinematix-be/src/services"
	"github.com/stretchr/testify/assert"
)

func TestListAllMovies(t *testing.T) {
	lib.DbConnection()

	// c := gin.Context{}

	req, _ := http.NewRequest("GET", "/movies", nil)
	// response := httptest.NewRecorder()
	// Router().ServeHTTP(response, req)

	// customerControllers.ListAllMovies(c, req)

	// if reflect.TypeOf(req) != reflect.TypeOf(&services.ResponseList{}) {
	// 	t.Error("failed test")
	// }
	assert.Equal(t, reflect.TypeOf(&services.ResponseList{}), reflect.TypeOf(req), "failed test")
}
