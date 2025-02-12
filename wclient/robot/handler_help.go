package robot

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient/model"
)

func helpHandler() {

	handlers["/help"] = &Handler{
		Level:    0,
		ChatAble: true,
		RoomAble: true,
		Describe: "查看帮助信息",
		Callback: func(msg *wcferry.WxMsg) string {
			user := args.GetMember(msg.Sender)
			// 生成指令菜单
			helper := []string{}
			for k, v := range handlers {
				if v.Level > 0 {
					if user == nil || v.Level > user.Level {
						continue // 没有权限
					}
				}
				if msg.IsGroup {
					if v.RoomAble { // 群聊指令
						helper = append(helper, k+" "+v.Describe)
					}
				} else {
					if v.ChatAble { // 私聊指令
						helper = append(helper, k+" "+v.Describe)
					}
				}
			}
			sort.Strings(helper)
			text := strings.Join(helper, "\n") + "\n"
			if user.Level > 0 {
				text += "级别 " + strconv.Itoa(user.Level) + "；"
			}
			if user.AiArgot != "" {
				text += "唤醒词 " + user.AiArgot + "；"
			}
			if len(args.LLM.Models) > 0 {
				text += "对话模型 " + user.GetModel().Family + "，"
				text += fmt.Sprintf("上下文长度 %d/%d", model.CountHistory(msg.Sender), args.LLM.HistoryNum) + "；"
			}
			if msg.IsGroup {
				room := args.GetChatRoom(msg.Roomid)
				if room.Level > 0 {
					text += "群级别 " + strconv.Itoa(room.Level) + "；"
					user := room.GetMember(msg.Sender)
					if user.Level > 0 {
						text += "群成员级别 " + strconv.Itoa(user.Level) + "；"
					}
				}
			}
			return text + "祝你好运！"
		},
	}

}
