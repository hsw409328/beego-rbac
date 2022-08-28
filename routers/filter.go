package routers

import (
	"beego-rbac/models"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/adapter/context"
	"github.com/beego/beego/v2/client/orm"
	"strings"
)

var (
	loginUrl                   = "/login"
	url404                     = "/404"
	filterExcludeURLMap        = make(map[string]int)
	filterOnlyLoginCheckURLMap = make(map[string]int)
)

func CheckRbac(uriStr string, ctx *context.Context) bool {
	tmpResult := ctx.Input.Session("LoginUserAllRBACResult")
	if tmpResult == nil {
		return false
	}
	returnResult := false
	for _, v := range tmpResult.([]orm.Params) {
		if v["name"].(string) == uriStr {
			returnResult = true
			break
		}
	}
	return returnResult
}

func CheckLogin(ctx *context.Context) (interface{}, bool) {
	userData := ctx.Input.Session("LoginUser")
	if userData == nil {
		return nil, false
	}
	return userData, true
}

func CheckRbacAll(uriStr string, ctx *context.Context) {
	info, err := CheckLogin(ctx)
	infoResult := info.(models.RbacUser)
	if !err {
		ctx.Redirect(302, loginUrl)
	}
	u := beego.AppConfig.String("superadmin")
	if infoResult.Username == u {
		return
	}
	err = CheckRbac(uriStr, ctx)
	if !err {
		ctx.Redirect(302, url404)
	}
}

var InitSetFilterUrl = func() {
	excludeUrl := beego.AppConfig.String("filterExcludeURL")
	if len(excludeUrl) > 0 {
		excludeUrlSlice := strings.Split(excludeUrl, ",")
		if len(excludeUrlSlice) > 0 {
			for _, v := range excludeUrlSlice {
				filterExcludeURLMap[v] = 1
			}
		}
	}
	checkLoginUrl := beego.AppConfig.String("filterOnlyLoginCheckURL")
	if len(checkLoginUrl) > 0 {
		checkLoginUrlSlice := strings.Split(checkLoginUrl, ",")
		if len(checkLoginUrlSlice) > 0 {
			for _, v := range checkLoginUrlSlice {
				filterOnlyLoginCheckURLMap[v] = 1
			}
		}
	}
}

var FilterRBAC = func(ctx *context.Context) {
	//判断URL是否排除
	if _, ok := filterExcludeURLMap[ctx.Request.URL.Path]; ok {
		return
	}
	_, okLogin := ctx.Input.Session("LoginUser").(models.RbacUser)
	//判断是否只验证登录的URL
	if _, ok := filterOnlyLoginCheckURLMap[ctx.Request.URL.Path]; okLogin && ok {
		return
	}
	if okLogin {
		CheckRbacAll(ctx.Request.URL.Path, ctx)
	} else {
		ctx.Redirect(302, loginUrl)
	}
}
