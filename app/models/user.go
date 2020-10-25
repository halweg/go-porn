package models

import (
    "github.com/astaxie/beego/orm"
    "time"
)

type User struct {
    Id int
    Name string
    Password string
    Status int
    AddTime int64
    Mobile string
    Avatar string
}

func init() {
    orm.RegisterModel(new (User))
}

//根据手机号判断用户是否存在
func IsUserMobile(mobile string) bool  {
    o := orm.NewOrm()
    user := User{Mobile:mobile}
    err := o.Read(&user, "Mobile")
    if err == orm.ErrNoRows {
        return false
    }
    if err == orm.ErrMissPK {
        return false
    }
    return true
}

func UserSave(mobile string, password string) error {
    o := orm.NewOrm()
    var user User
    user.Name = ""
    user.Mobile = mobile
    user.Password = password
    user.Status = 1
    user.AddTime = time.Now().Unix()
    _, err := o.Insert(&user)
    return err
}

