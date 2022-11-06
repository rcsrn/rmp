package database

type Builder struct {
	rolas []*Rola
	dbPath string
}

func CreateNewBuilder(rolas []*Rola, dbPath string) *Builder {
	return &Builder{rolas, dbPath}
}

func (builder *Builder) BuildDataBase() (*DataBase, error) {
	database, err := CreateNewDataBase()
	if err != nil {
		return nil, err
	}
	return database, nil
}

