package model

type Page struct {
	Count    int64       `json:"count"`
	PageNo   int         `json:"pageNo" form:"pageNo"`
	PageSize int         `json:"pageSize" form:"pageSize"`
	Data     interface{} `json:"data"`
}

func NewPage(pageNo int, pageSize int) *Page {
	return &Page{PageNo: pageNo, PageSize: pageSize}
}

func (a *Page) SetData(data interface{}) *Page {
	a.Data = data
	return a
}
