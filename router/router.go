package router

import (
	"app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Rou struct {
	a *usecase.K
}

func Init(r *usecase.K) *Rou {
	return &Rou{a: r}
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

// @Summary Print all elements
// @Description Retrieve a hello message from the server
// @Produce json
// @Success 200 {object} Print
// @Router /print [get]
func (w *Rou) Print(ctx *gin.Context) {
	e := w.a.Print()
	ctx.JSON(http.StatusOK, Print{
		Print: e,
	})
}

// @Summary Save the data
// @Description Insert data at the start
// @Produce json
// @Param key path string true "Key parameter"
// @Success 200 {object} HelloResponse
// @Router /start/{key} [post]
func (w *Rou) Insertatstarting(ctx *gin.Context) {
	s := ctx.Param("key")
	w.a.Insertatstarting(s)
	ctx.JSON(http.StatusOK, HelloResponse{
		Mess: "added successfully",
	})
}

// @Summary Length of linked list
// @Description Retrieve a length of the linked list from the server
// @Produce json
// @Success 200 {object} Length
// @Router /length [get]
func (w *Rou) Length(ctx *gin.Context) {
	e := w.a.Length()
	ctx.JSON(http.StatusOK, Length{
		Length: e,
	})
}

// @Summary Save the data at the end of the linked list
// @Description Insert data at the end
// @Produce json
// @Param key path string true "Key parameter"
// @Success 200 {object} HelloResponse
// @Router /end/{key} [post]
func (w *Rou) Insertatending(ctx *gin.Context) {
	s := ctx.Param("key")
	w.a.InsertAtEnding(s)
	ctx.JSON(http.StatusOK, HelloResponse{
		Mess: "added successfully",
	})
}

// @Summary delete at begining
// @Description deleting the data in the at begining
// @Produce json
// @Success 200 {object} HelloResponse
// @Router /front [delete]
func (w *Rou) DelAtBeg(ctx *gin.Context) {

	s := w.a.DelAtBeg()
	ctx.JSON(http.StatusOK, HelloResponse{
		Mess: s,
	})
}

// @Summary delete at ending
// @Description deleting the data in the ending
// @Produce json
// @Success 200 {object} HelloResponse
// @Router /end [delete]
func (w *Rou) DelAtEnd(ctx *gin.Context) {

	s := w.a.DelAtEnd()
	ctx.JSON(http.StatusOK, HelloResponse{
		Mess: s,
	})
}
