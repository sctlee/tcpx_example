package chatroom

import (
	. "features/chatroom/action"

	"github.com/sctlee/tcpx"
)

//TODO: redefine struct function, then move the usage from here to example
var chatroomAction = NewChatroomAction()

var Router = map[string]tcpx.RouteFun{
	"list": chatroomAction.List,
	"view": chatroomAction.View,
	"join": chatroomAction.Join,
	"exit": chatroomAction.Exit,
	"send": chatroomAction.Send,
}

// f := auth.PermissionDecorate(client, chatroomAction.Send, auth.IsLogin)
// responseMsg = f(client, params)
