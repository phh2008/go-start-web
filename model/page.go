package model

type Page struct {
	Count    int64
	PageNo   int
	PageSize int
	Data     interface{}
}

func NewPage(pageNo int, pageSize int) *Page {
	return &Page{PageNo: pageNo, PageSize: pageSize}
}

func (a *Page) SetData(data interface{}) *Page {
	a.Data = data
	return a
}
