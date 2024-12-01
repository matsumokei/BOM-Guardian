package server

import (
	"encoding/json"
	"net/http"

	"github.com/matsumokei/BOM-Guardian/pkg/db"
	"github.com/matsumokei/BOM-Guardian/pkg/utils"
)



type ServerInterface interface {
	HandleSearchPackages(w http.ResponseWriter, r *http.Request)
	//HandleImportBOM(w http.ResponseWriter, r *http.Request)
	//HandleDeleteComponent(w http.ResponseWriter, r *http.Request)
}

type serverInterface struct{
	bs repository.BomStore
}

func NewServerInterface(bs repository.BomStore) ServerInterface {
	return &serverInterface{bs}
}


// func NewServer(tr repository.TodoRepository) TodoController {
// 	return &todoController{tr}
// }

// (GET /analysis/dependencies)
func (si *serverInterface) HandleSearchPackages(w http.ResponseWriter, r *http.Request) {
	pkgs, err := si.bs.GetPkgs()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// var todoResponses []dto.TodoResponse
	// for _, v := range todos {
	// 	todoResponses = append(todoResponses, dto.TodoResponse{Id: v.Id, Title: v.Title, Content: v.Content})
	// }

	// var todosResponse dto.TodosResponse
	// todosResponse.Todos = todoResponses

	output, _ := json.MarshalIndent(pkgs, "", "\t\t")
	utils.WriteJSON(w, http.StatusOK, pkgs)
	w.Write(output)
}

// func (si serverInterface) HandleImportBOM(w http.ResponseWriter, r *http.Request) {

// }


