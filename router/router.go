package router

import (
	"app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	a *usecase.Usecase
}

// Init initializes a new Router instance with a usecase.K instance
func Init(r *usecase.Usecase) *Router {
	return &Router{a: r}
}

type HelloResponse struct {
	Mess string `json:"mess"`
}

type Print struct {
	Print string `json:"print"`
}

type Length struct {
	Length int `json:"length"`
}

// Print handles GET requests to retrieve all elements in the linked list
// @Summary Retrieve all elements
// @Description Retrieve all elements in the linked list and return them as a concatenated string
// @Produce json
// @Success 200 {object} Print
// @Router /print [get]
func (w *Router) Print(ctx *gin.Context) {
	message := w.a.Print()
	ctx.JSON(http.StatusOK, Print{
		Print: message,
	})
}

// Insertatstarting handles POST requests to insert data at the start of the linked list
// @Summary Insert data at the start
// @Description Insert a new element at the beginning of the linked list
// @Produce json
// @Param key path string true "Key parameter"
// @Success 200 {object} HelloResponse
// @Router /start/{key} [post]
func (w *Router) Insertatstarting(ctx *gin.Context) {
	data := ctx.Param("key")
	w.a.Insertatstarting(data)
	ctx.JSON(http.StatusOK, HelloResponse{
		Mess: "added successfully",
	})
}

// Length handles GET requests to retrieve the length of the linked list
// @Summary Retrieve the length of the linked list
// @Description Retrieve and return the number of elements in the linked list
// @Produce json
// @Success 200 {object} Length
// @Router /length [get]
func (w *Router) Length(ctx *gin.Context) {
	length := w.a.Length()
	ctx.JSON(http.StatusOK, Length{
		Length: length,
	})
}

// Insertatending handles POST requests to insert data at the end of the linked list
// @Summary Insert data at the end
// @Description Insert a new element at the end of the linked list
// @Produce json
// @Param key path string true "Key parameter"
// @Success 200 {object} HelloResponse
// @Router /end/{key} [post]
func (w *Router) Insertatending(ctx *gin.Context) {
	data := ctx.Param("key")
	w.a.InsertAtEnding(data)
	ctx.JSON(http.StatusOK, HelloResponse{
		Mess: "added successfully",
	})
}

// DelAtBeg handles DELETE requests to delete the first element in the linked list
// @Summary Delete the first element
// @Description Delete the first element in the linked list
// @Produce json
// @Success 200 {object} HelloResponse
// @Router /front [delete]
func (w *Router) DelAtBeg(ctx *gin.Context) {
	message := w.a.DelAtBeg()
	ctx.JSON(http.StatusOK, HelloResponse{
		Mess: message,
	})
}

// DelAtEnd handles DELETE requests to delete the last element in the linked list
// @Summary Delete the last element
// @Description Delete the last element in the linked list
// @Produce json
// @Success 200 {object} HelloResponse
// @Router /end [delete]
func (w *Router) DelAtEnd(ctx *gin.Context) {
	message := w.a.DelAtEnd()
	ctx.JSON(http.StatusOK, HelloResponse{
		Mess: message,
	})
}
