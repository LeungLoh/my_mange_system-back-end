package api

import (
	"fmt"
	"my_mange_system/common"
	"my_mange_system/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserLoginParams struct {
	Username string `form:"username"`
	Password string `form:"password"`
	City     string `form:"city"`
}

type UserListParams struct {
	Username string `form:"username"`
	Roleid   int    `form:"roleid"`
	Offset   int    `form:"offset"`
	Limit    int    `form:"limit"`
}

type UserListHandle struct {
	Username string `form:"username"`
	Password string `form:"password"`
	UserId   int    `form:"userId"`
}

func UserRegister(ctx *gin.Context) {

}

func UserLogin(ctx *gin.Context) {
	var userloginparams UserLoginParams
	var res common.Result
	if ctx.ShouldBind(&userloginparams) == nil {
		if server.CheckOutUser(ctx, userloginparams.Username, userloginparams.Password) == true {
			server.UpdateLoginInfo(userloginparams.City, userloginparams.Username)
			res = common.Result{Httpcode: http.StatusOK, Msg: "登录成功"}
		} else {
			res = common.Result{Httpcode: http.StatusBadRequest, Msg: "账号密码错误"}
		}
	} else {
		res = common.Result{Httpcode: http.StatusBadRequest, Msg: "用户数据解析失败"}
	}
	ctx.Set("Res", res)
	ctx.Next()
}

func UserInfo(ctx *gin.Context) {
	var userinfoparams UserLoginParams
	ctx.ShouldBindQuery(&userinfoparams)
	username, roleid, city, lastlogintime := server.GetUserinfo(userinfoparams.Username)
	res := common.Result{Httpcode: http.StatusOK, Msg: "获取信息成功", Data: gin.H{"username": username, "roleid": roleid, "city": city, "lastlogintime": lastlogintime}}
	ctx.Set("Res", res)
	ctx.Next()
}

func UserList(ctx *gin.Context) {
	var userlistparams = UserListParams{
		Username: "",
		Roleid:   0,
		Offset:   1,
		Limit:    10,
	}
	ctx.ShouldBindQuery(&userlistparams)
	users, total := server.GetUsetList(userlistparams.Username, userlistparams.Roleid, (userlistparams.Offset-1)*userlistparams.Limit, userlistparams.Limit)
	res := common.Result{Httpcode: http.StatusOK, Msg: "获取信息成功", Data: gin.H{"users": users, "total": total}}
	ctx.Set("Res", res)
	ctx.Next()
}

func UserDelete(ctx *gin.Context) {
	// var user UserListHandle

	user := common.GetSession(ctx, "user")
	if user != nil {
		fmt.Println(user)
	} else {
		fmt.Println("222222222222222222")
	}
	res := common.Result{Httpcode: 200, Msg: "非管理员无法删除"}
	ctx.Set("Res", res)
	ctx.Next()
	// fmt.Println(user)
	// if user["roleid"] == 1 {
	// 	res := common.Result{Httpcode: http.StatusUnauthorized, Msg: "非管理员无法删除"}
	// 	ctx.Set("Res", res)
	// 	ctx.Next()
	// }
	// if ctx.ShouldBind(&userloginparams) == nil{

	// }
}

func UserLogout(ctx *gin.Context) {

}
