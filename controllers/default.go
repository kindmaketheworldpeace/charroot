package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
	"net/http"
	  "github.com/gorilla/websocket"
)

type MainController struct {
	beego.Controller
}
type EntryController struct {
	beego.Controller
}
type ChatRoomController struct {
	beego.Controller
}
func (c *MainController) Get() {

	c.TplName = "index.html"
}

func (c *EntryController) Post() {
             account := make(map[string]string)
            if err := json.Unmarshal(c.Ctx.Input.RequestBody, &account); err == nil {
                    user_info:=UserInfo{
                    Username:account["username"],
                    Message: account["username"]+" coming!",
                    Type: 1,
                    }
                     coming=append(coming,user_info)

            }else {
                fmt.Println(err)
            }
            c.ServeJSON()

}
var (
    upgrader = websocket.Upgrader {
        // 读取存储空间大小
        ReadBufferSize:1024,
        // 写入存储空间大小
        WriteBufferSize:1024,
        // 允许跨域
        CheckOrigin: func(r *http.Request) bool {
            return true
        },
    }
)
func (c *ChatRoomController) Get() {
    ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)

    if err != nil {
        fmt.Println(err)

    }
    coming :=coming[0:len(coming)]
    userInfo := coming[0]

     ch :=make(chan UserInfo)
     go sendMessage(ws,ch)
     ch<-userInfo
     userInfo.Message = userInfo.Username + " has arrived"
     userInfo.Type =2
     message<-userInfo
     entry<-ch

    for {
		_, p, err := ws.ReadMessage()

		if err != nil {
		fmt.Println(err)
		    break
		}
		userInfo.Message = string(p)
		message<-userInfo

	}

	leave<-ch
	userInfo.Message = userInfo.Username + " has left"
	userInfo.Type =3
    message<-userInfo
    ws.Close()
    c.ServeJSON()

}
func sendMessage(ws *websocket.Conn,  ch  <-chan  UserInfo) {
    for msg := range ch {
       fmt.Println(msg.Message)
      s,_:=json.Marshal(msg)
      fmt.Println(string(s))
      ws.WriteMessage(websocket.TextMessage, []byte(s))
    }

}