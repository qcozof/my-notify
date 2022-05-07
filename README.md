# my-notify
system:
  proxy-address: 'socks5://127.0.0.1:9081'

pushplus:
  api-url: "http://www.pushplus.plus/send"
  token: ""
  enable: false

telegram:
  api-url: "https://api.telegram.org"
  token: "1234:xxxxx"
  chat-id: "-100123456"
  enable: true

discord:
  token: "Bot xx.yy.zz"
  enable: true

email:
  host: "smtp.exmail.qq.com"
  port: 465
  username: "xxx@qq.cn"
  password: "123456"
  enable: true

#短信、slider滑块验证
sms:
  access-key-id: "xxx"
  access-key-secret: "xxx"
  template-code: "SMS_222222"
  sign-name-json: '["xx平台"]'
  slider-scene: ''
  slider-app-key: ''
  enable: true

#短信－中国移动
china-mobile-sms:
  api-url: ""
  ec-name: ''
  ap-id: ""
  secret-key: ""
  template-id: ""
  sign: "xxx"
  add-serial: ""

#短信－中国联通
ums:
  url: "https://api.ums86.com:9600/sms/Api/Send.do"
  spcode: '123456'
  loginname: 'aaaa'
  password: 'xxxx'

dingding:
  api-url: "https://oapi.dingtalk.com/robot/send?access_token=xxxxx"
  enable: true
