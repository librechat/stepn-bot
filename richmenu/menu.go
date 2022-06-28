package richmenu

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func RichMenu() *linebot.RichMenu {
	return &linebot.RichMenu{
		Size: linebot.RichMenuSize{
			Width:  800,
			Height: 270,
		},
		Selected:    true,
		Name:        "Exchange Menu",
		ChatBarText: "快速匯率",
		Areas: []linebot.AreaDetail{
			{
				Bounds: linebot.RichMenuBounds{
					X:      0,
					Y:      0,
					Width:  200,
					Height: 270,
				},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypePostback,
					Data: "price gst 100",
				},
			},
			{
				Bounds: linebot.RichMenuBounds{
					X:      200,
					Y:      0,
					Width:  200,
					Height: 270,
				},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypePostback,
					Data: "price gmt",
				},
			},
			{
				Bounds: linebot.RichMenuBounds{
					X:      400,
					Y:      0,
					Width:  200,
					Height: 270,
				},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypePostback,
					Data: "price sol",
				},
			},
			{
				Bounds: linebot.RichMenuBounds{
					X:      600,
					Y:      0,
					Width:  200,
					Height: 270,
				},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypePostback,
					Data: "price usd",
				},
			},
		},
	}
}
