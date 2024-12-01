package repository

import (

	_ "github.com/lib/pq"

)


type BomStore interface {
	GetPkgs() (pkg []*PackageEntity, err error)
	//InsertComps(todo Bom) (id int, err error)
	//UpdateComps(todo Bom) (err error)
	//DeleteComps(id int) (err error)
}

type bomStore struct {
}

func NewTodoRepository() BomStore {
	return &bomStore{}
}

func (bs *bomStore) GetPkgs() ([]*PackageEntity, error){

	rows, err := Db.Query("SELECT id, name, version, purl FROM bomstore ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	pkgs := make([]*PackageEntity, 0)
	for rows.Next() {
		pkg := &PackageEntity{}
		//ToDo: Create function scanRowsIntoPacakge(rows)
		err = rows.Scan(&pkg.Id, &pkg.Name, &pkg.Version, &pkg.Purl)
		if err != nil {
			return nil, err
		}
		pkgs = append(pkgs, pkg)
	}

	return pkgs, nil
}