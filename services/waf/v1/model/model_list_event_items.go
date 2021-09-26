package model

import (
	"encoding/json"

	"strings"
)

// 攻击类型
type ListEventItems struct {
	// 事件id

	Id *string `json:"id,omitempty"`
	// 次数

	Time *int64 `json:"time,omitempty"`
	// 策略id

	Policyid *string `json:"policyid,omitempty"`
	// 源ip

	Sip *string `json:"sip,omitempty"`
	// 域名

	Host *string `json:"host,omitempty"`
	// 攻击的url链接

	Url *string `json:"url,omitempty"`
	// 攻击类型（XSS攻击：xss，sqli，命令注入：cmdi，恶意爬虫：robot，本地文件包含：lfi，远程文件包含：rfi，网站木马：webshell，cc攻击：cc，精准防护：custom_custom，IP黑白名单：custom_whiteblackip，地理位置访问控制：custom_geoip，防篡改：antitamper，反爬虫：anticrawler，网站信息防泄漏：leakage，非法请求：illegal，其它类型攻击：vuln）

	Attack *string `json:"attack,omitempty"`
	// 命中的规则id

	Rule *string `json:"rule,omitempty"`
	// 命中的载荷

	Payload *string `json:"payload,omitempty"`
	// 防护动作

	Action *string `json:"action,omitempty"`
	// 请求方法和路径

	RequestLine *string `json:"request_line,omitempty"`

	Headers *ListEventItemsHeaders `json:"headers,omitempty"`
	// 请求cookie

	Cookie *string `json:"cookie,omitempty"`
	// 响应码状态

	Status *string `json:"status,omitempty"`
	// 区域

	Region *string `json:"region,omitempty"`
	// 域名id

	HostId *string `json:"host_id,omitempty"`
	// 响应时长

	ResponseTime *int64 `json:"response_time,omitempty"`
	// 响应体大小

	ResponseSize *int32 `json:"response_size,omitempty"`
	// 响应体

	ResponseBody *string `json:"response_body,omitempty"`
}

func (o ListEventItems) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEventItems struct{}"
	}

	return strings.Join([]string{"ListEventItems", string(data)}, " ")
}
