package obj

type Pagination struct {
	Count int
	PageSize int
	PageNum int `form:"page"`
	Data []interface{}
}