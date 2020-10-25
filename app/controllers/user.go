package controllers

import (
    "app/models"
    "github.com/astaxie/beego"
    "regexp"
)

type UserController struct {
    beego.Controller
}

// @router /register/save [post]
func (this *UserController) SaveRegister() {

    mobile := this.GetString("mobile")
    password := this.GetString("password")

    if mobile == "" {
        this.Data["json"] = ReturnError(4001, "手机号不能为空")
        this.ServeJSON()
    }

    if isMobileOk, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile); !isMobileOk {
        this.Data["json"] = ReturnError(4002, "手机号码格式不对！")
        this.ServeJSON()
    }

    if password == "" {
        this.Data["json"] = ReturnError(4001, "密码不能为空")
        this.ServeJSON()
    }

    if status := models.IsUserMobile(mobile); status {
        this.Data["json"] = ReturnError(4005, "此手机号已经被注册!")
        this.ServeJSON()
    }

    err := models.UserSave(mobile, MD5V(password))
    if err != nil {
        this.Data["json"] = ReturnError(5000, err.Error())
        this.ServeJSON()

    }

    this.Data["json"] = ReturnSuccess(0, "注册成功!", nil, 0)
    this.ServeJSON()
}