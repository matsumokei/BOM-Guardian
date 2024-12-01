package server

import (
	"net/http/httptest"
	"testing"
	repository "github.com/matsumokei/BOM-Guardian/pkg/db"
)


func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/", nil)

	bomStore := &mockBomStore{}
	handler := NewServerInterface(bomStore)

	handler.HandleSearchPackages(w, r)

	if w.Code != 200 {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type is %v", w.Header().Get("Content-Type"))
	}

	// body := make([]byte, w.Body.Len())
	// w.Body.Read(body)
	// var todosResponse dto.TodosResponse
	// json.Unmarshal(body, &todosResponse)
	// if len(todosResponse.Todos) != 0 {
	// 	t.Errorf("Response is %v", todosResponse.Todos)
	// }

}


type mockBomStore struct{}

func (m *mockBomStore) GetPkgs() ([]*repository.PackageEntity, error) {
	return []*repository.PackageEntity{}, nil
}

