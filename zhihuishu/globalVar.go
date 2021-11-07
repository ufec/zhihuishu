package zhihuishu

var header = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Encoding": "gzip, deflate, br",
	"Accept-Language": "zh-CN,zh;q=0.9",
	"Connection":      "keep-alive",
	"Content-Type":    "application/x-www-form-urlencoded",
	"Host":            "onlineservice.zhihuishu.com",
	"Origin":          "https://onlineweb.zhihuishu.com",
	"Referer":         "https://onlineweb.zhihuishu.com/",
	"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36",
}

var (
	nowTime = 0
	cookie1 = ""
	cookie2 = ""
	uuid    = ""
	isNeed  int
)
