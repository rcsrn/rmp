package database

type Builder struct {
	rolas []*Rola
}

func CreateNewBuilder() *Builder {
	return &Builder{}
}

func (builder *Builder) SetRolas(rolas []*Rola) {
	builder.rolas = rolas
}

func (builder *Builder) BuildDataBase() *DataBase {
	database := CreateNewDataBase()
	for _, rola := range(builder.rolas) {
		
	}
	return nil
}
