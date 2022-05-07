/**
 * @Description telegram notify
 * @Author $
 * @Date $ $
 **/
package notify

import (
	"encoding/json"
	"fmt"
	"github.com/qcozof/my-notify/global"
	"net/http"
	"net/url"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func initDiscord(authorization string) {
	if global.DISCORD != nil {
		return
	}

	//初始化
	discord, err := discordgo.New(authorization)
	if err != nil {
		fmt.Println("初始化出错：", err) //按回车退出
	}

	//设置代理
	proxyAddress := global.SERVER_CONFIG.SystemConfig.ProxyAddress
	if strings.TrimSpace(proxyAddress) != ""{
		proxyUrl, _ := url.Parse(proxyAddress)
		discord.Client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	}


	global.DISCORD = discord

}
func Discord(title,content string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Discord err:", err)
		}
	}()

	var dcConfig = global.SERVER_CONFIG.DiscordConfig
	var mpChannelIds map[string]string
	json.Unmarshal([]byte(dcConfig.ChannelIdJson),&mpChannelIds)
	channelId := mpChannelIds[global.IN_ARG]
	if len(channelId)==0{
		fmt.Println("Discord channelId不存在！")
		return
	}

	initDiscord(dcConfig.Token)

	if !dcConfig.Enable {
		fmt.Println("discord not enable !")
		return
	}

	var msgEmbed = discordgo.MessageEmbed{
		Title: title,
		Description:content,
	}
	msg, err :=global.DISCORD.ChannelMessageSendEmbed(channelId,&msgEmbed)
	//msg, err := global.DISCORD.ChannelMessageSend(channelId, content)
	//fmt.Println(msg, err)

	if err == nil {
		fmt.Println(fmt.Sprintf("discord 消息已发送！chat_id:%s", msg.ID))
		return
	}
	fmt.Println(fmt.Sprintf("discord 消息发送失败%s", err))

}
