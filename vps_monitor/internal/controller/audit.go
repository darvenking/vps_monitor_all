package controller

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"vps_monitor/internal/model"
	"vps_monitor/internal/param"
	"vps_monitor/utility/res"
)

// AuditList 审核列表
func AuditList(r *ghttp.Request) {
	var p param.PageParam
	err := r.Parse(&p)
	if err != nil {
		res.Fail(r, "ok")
	}
	var sub []model.SubmitSite
	db := model.GetSubmitSiteDB()
	if p.Status != 0 {
		db.Where("status = ?", p.Status)
	}
	db.Limit(p.Size).Offset((p.Page - 1) * p.Size).Order("status asc,create_at desc").Find(&sub)
	var count int64
	db.Count(&count)
	res.Success(r, g.Map{
		"total": count,
		"data":  sub,
	})
}

// Audit 审核
func Audit(r *ghttp.Request) {
	var p param.AuditSite
	err := r.Parse(&p)
	if err != nil {
		res.Fail(r, "ok")
	}
	var sub model.SubmitSite
	model.GetSubmitSiteDB().First(&sub, p.Id)
	m := &model.SiteConfig{
		URL:         sub.URL,
		NoStockFlag: p.NoStockFlag,
		PriceFlag:   p.PriceFlag,
		NameFlag:    p.NameFlag,
		Cookies:     p.Cookies,
	}
	model.GetSubmitSiteDB().Where("id = ?", p.Id).Update("status", 2)
	model.GetSiteConfigDB().Save(m)
	res.Success(r, "ok")
}
