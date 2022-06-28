package richmenu

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func CreateRichMenu(bot *linebot.Client) error {
	menu, image := RichMenu()
	resp, err := bot.CreateRichMenu(*menu).Do()
	if err != nil {
		return err
	}

	if _, err := bot.UploadRichMenuImage(resp.RichMenuID, image).Do(); err != nil {
		return err
	}

	if _, err := bot.SetDefaultRichMenu(resp.RichMenuID).Do(); err != nil {
		return err
	}
	return nil
}

func RichMenu() (*linebot.RichMenu, string) {
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
					Data: "prices gst 100",
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
					Data: "prices gmt",
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
					Data: "prices sol",
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
					Data: "prices usd",
				},
			},
		},
	}, "./richmenu/images/menu.png"
}
