package controllers
import (

	//"github.com/astaxie/beego"
	//"github.com/gorilla/websocket"
	//"fmt"

)
type UserInfo struct {
    Username string
    Message string
    Type   int

}
type client chan<- UserInfo
var message = make(chan UserInfo)
var entry = make(chan client)
var leave = make(chan client)
var coming = make([]UserInfo ,0)



func chatroom (Users map[client]bool){

    for {
    select {
    case msg := <-message:
        for user := range Users {
            user<-msg
        }
    case  cli:=<-entry:
        Users[cli] = true
      case cli := <- leave:
        delete(Users,cli)
        close(cli)
        }
    }
}

func init () {
    Users := make(map[client]bool)
    go chatroom(Users)
}