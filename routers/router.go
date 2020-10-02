package routers

import (
	"chat/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
     beego.Router("/entry", &controllers.EntryController{})
      beego.Router("/chatroom", &controllers.ChatRoomController{})
}
