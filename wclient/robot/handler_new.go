package robot

import (
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient/model"
)

func newHandler() {

	handlers["/new"] = &Handler{
		Level:    0,
		ChatAble: true,
		RoomAble: true,
		Describe: "重置上下文内容",
		Callback: func(msg *wcferry.WxMsg) string {
			model.ResetHistory(msg.Sender)
			return "已重置上下文"
		},
	}

}
