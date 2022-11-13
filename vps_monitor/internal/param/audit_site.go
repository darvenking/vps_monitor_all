package param

type AuditSite struct {
	Id          uint   `json:"id" v:"required#id不能为空"`
	NameFlag    string `json:"nameFlag" v:"required#请输入商品名选择器"`
	PriceFlag   string `json:"priceFlag" v:"required#请输入价格选择器"`
	NoStockFlag string `json:"noStockFlag"`
	Cookies     string `json:"cookies"`
}
