package routes

import (
	"extensao-api/controller"
	"net/http"
)

var booksRoutes = []Routes{
	{
		URI:    "/books/search",
		Method: http.MethodPost,
		Func:   controller.HandleSearch,
		Auth:   true,
	},
	{
		URI:    "/books",
		Method: http.MethodPost,
		Func:   controller.CreateBook,
		Auth:   true,
	},
	{
		URI:    "/books",
		Method: http.MethodGet,
		Func:   controller.FetchBooks,
		Auth:   true,
	},
	{
		URI:    "/books/{bookID}",
		Method: http.MethodPut,
		Func:   controller.UpdateBook,
		Auth:   true,
	},
	{
		URI:    "/books/{bookID}",
		Method: http.MethodDelete,
		Func:   controller.DeleteBook,
		Auth:   true,
	},
}
