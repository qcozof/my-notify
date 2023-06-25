package config

type ServerConfig struct {
	SystemConfig         SystemConfig         `mapstructure:"system" json:"system" yaml:"system"`
	PushPlusConfig       PushPlusConfig       `mapstructure:"pushplus" json:"pushplus" yaml:"pushplus"`
	TelegramConfig       TelegramConfig       `mapstructure:"telegram" json:"telegram" yaml:"telegram"`
	DiscordConfig        DiscordConfig        `mapstructure:"discord" json:"discord" yaml:"discord"`
	EmailConfig          EmailConfig          `mapstructure:"email" json:"email" yaml:"email"`
	SmsConfig            SmsConfig            `mapstructure:"sms" json:"sms" yaml:"sms"`
	DingDingConfig       DingDingConfig       `mapstructure:"dingding" json:"dingding" yaml:"dingding"`
	ChinaMobileSmsConfig ChinaMobileSmsConfig `mapstructure:"china-mobile-sms" json:"chinaMobileSms" yaml:"china-mobile-sms"`
	UmsConfig            UmsConfig            `mapstructure:"ums" json:"ums" yaml:"ums"`
	ZapConfig            ZapConfig            `mapstructure:"zap" json:"zap" yaml:"zap"`
}

type SystemConfig struct {
	IsDebug      bool   `mapstructure:"is-debug" json:"isDebug" yaml:"is-debug" `
	ProxyAddress string `mapstructure:"proxy-address" json:"proxyAddress" yaml:"proxy-address" `
}

type PushPlusConfig struct {
	ApiUrl string `mapstructure:"api-url" json:"ApiUrl" yaml:"api-url" `
	Token  string `mapstructure:"token" json:"token" yaml:"token" `
	Enable bool   `mapstructure:"enable" json:"enable" yaml:"enable" `
}

type TelegramConfig struct {
	ApiUrl string `mapstructure:"api-url" json:"apiUrl" yaml:"api-url" `
	Token  string `mapstructure:"token" json:"token" yaml:"token" `
	ChatId string `mapstructure:"chat-id" json:"chatId" yaml:"chat-id" `
	Enable bool   `mapstructure:"enable" json:"enable" yaml:"enable" `
}

type DiscordConfig struct {
	Token         string `mapstructure:"token" json:"token" yaml:"token" `
	ChannelIdJson string `mapstructure:"channel-id-json" json:"channelIdJson" yaml:"channel-id-json" `
	Enable        bool   `mapstructure:"enable" json:"enable" yaml:"enable" `
}

type EmailConfig struct {
	Host     string   `mapstructure:"host" json:"host" yaml:"host"`
	Port     int      `mapstructure:"port" json:"port" yaml:"port"`
	Username string   `mapstructure:"username" json:"username" yaml:"username"`
	Password string   `mapstructure:"password" json:"password" yaml:"password"`
	EmailTo  []string `mapstructure:"email-to" json:"EmailTo" yaml:"email-to"`
	Enable   bool     `mapstructure:"enable" json:"enable" yaml:"enable" `
}

type SmsConfig struct {
	AccessKeyId     string `mapstructure:"access-key-id" json:"accessKeyId" yaml:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret" json:"accessKeySecret" yaml:"access-key-secret"`
	TemplateCode    string `mapstructure:"template-code" json:"templateCode" yaml:"template-code"`
	SignNameJson    string `mapstructure:"sign-name-json" json:"signNameJson" yaml:"sign-name-json"`
	SliderScene     string `mapstructure:"slider-scene" json:"sliderScene" yaml:"slider-scene"`
	SliderAppKey    string `mapstructure:"slider-app-key" json:"sliderAppKey" yaml:"slider-app-key"`
	Enable          bool   `mapstructure:"enable" json:"enable" yaml:"enable" `
}

// 移动短信
type ChinaMobileSmsConfig struct {
	ApiUrl     string `mapstructure:"api-url" json:"apiUrl" yaml:"api-url"`
	EcName     string `mapstructure:"ec-name" json:"ecName" yaml:"ec-name"`
	ApId       string `mapstructure:"ap-id" json:"accessKeySecret" yaml:"ap-id"`
	SecretKey  string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
	TemplateId string `mapstructure:"template-id" json:"templateId" yaml:"template-id"`
	Sign       string `mapstructure:"sign" json:"sign" yaml:"sign"`
	AddSerial  string `mapstructure:"add-serial" json:"addSerial" yaml:"add-serial"`
	ToMobiles  string `mapstructure:"to-mobiles" json:"to-mobiles" yaml:"to-mobiles"`
	Enable     bool   `mapstructure:"enable" json:"enable" yaml:"enable" `
}

// 联通短信
type UmsConfig struct {
	Url          string `mapstructure:"url" json:"url" yaml:"url"`
	TemplateCode string `mapstructure:"template-code" json:"templateCode" yaml:"template-code"`
	SPCode       string `mapstructure:"spcode" json:"spcode" yaml:"spcode"`
	LoginName    string `mapstructure:"loginname" json:"loginname" yaml:"loginname"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	ToMobiles    string `mapstructure:"to-mobiles" json:"to-mobiles" yaml:"to-mobiles"`
	Enable       bool   `mapstructure:"enable" json:"enable" yaml:"enable" `
}

type DingDingConfig struct {
	ApiUrl string `mapstructure:"api-url" json:"apiUrl" yaml:"api-url" `
	Enable bool   `mapstructure:"enable" json:"enable" yaml:"enable" `
}

type ZapConfig struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Directory     string `mapstructure:"directory" json:"directory"  yaml:"directory"`
	LinkName      string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"`
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`
}
