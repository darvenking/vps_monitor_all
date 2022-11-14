package param

type PageParam struct {
	Status int // 分页号码
	Page   int `d:1  v:"min:0#分页号码错误"`     // 分页号码
	Size   int `d:10 v:"max:50#分页数量最大50条"` // 分页数量，最大50
}
