package controllers
import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/go-bookstore/pkg/utils"
	"github.com/Harishgouda06022004/go-bookstore/pkg/models"
)
var NewBook models.Book
func GetBook(w http.ResponseWriter,r *http.Request){
	newBooks:=models.GetALLBooks()
	res,_:=json.Marshal(newBooks)
	w.Header().set("Cotent-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookById(w http.ResponseWriter,r *http.Request){
	vars:=mux.vars(r)
	bookId:=vars["bookId"]
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("error while parsing")
	}
	bookDetails,_:=models.GetBookByID(ID)
	res,_:=json.Marshal(bookDetails)
	w.Header().set("content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter,r *http.Request){
	CreateBook:=&models.Book{}
	utils.ParseBody(r,CreateBook)
	b:=CreateBook.CreateBook()
	res,_:=json.Marshal(b)
	w.Header().set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter,r *http.Request){
	vars:=mux.vars(r)
	bookId:=vars["bookId"]
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("error while parsing")
	}
	book:=model.DeleteBook(ID)
	res,_:=json.Marshal(book)
	w.Header().set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateBook(w http.ResponseWriter,r *http.Request){
	var updateBook:=&models.Book{}
	utils.ParseBody(r,updateBook)
	vars:=mux.vars(r)
	bookId:=vars["bookId"]
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("error while parsing")
	}
	bookDetails,db:=models.GetBookBYId(ID)
	if updateBook.Name!=""{
		bookDetails.Name:=updateBook.Name
	}
	if updateBook.Author!=""{
		bookDetails.Author:=updateBook.Author
	}
	if updateBook.Publication!=""{
		bookDetails.Publication:=updateBook.Publication
	}
	db.Save(&bookDetails)
	res,_:=json.Marshal(bookDetails)
	w.Header().set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
