// Package arrayvsslice handles application configuration, including loading environment variables.
package arrayvsslice

import (
	"app/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArrayVsSlice struct{}

func (e ArrayVsSlice) ModifyASlice(c *gin.Context) {
	a := []int{1, 2, 3000}
	b := a
	b[0] = 100

	fmt.Println("a:", a, "b:", b)
	data := map[string]interface{}{"a": a, "b": b}

	helpers.RenderJSON(c, http.StatusOK, data, "Success", "no error", true)
}

func (e ArrayVsSlice) AppendToASlice(c *gin.Context) {
	a := []int{1, 2, 3}
	b := a
	b = append(b, 4)
	b[0] = 100

	fmt.Println("a:", a, "b:", b)
	data := map[string]interface{}{"a": a, "b": b}

	helpers.RenderJSON(c, http.StatusOK, data, "Success", "no error", true)
}
