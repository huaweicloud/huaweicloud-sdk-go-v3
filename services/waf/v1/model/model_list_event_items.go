package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

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
	// 攻击类型:   - vuln：其它攻击类型   - sqli： sql注入攻击   - lfi： 本地文件包含  - cmdi：命令注入攻击   - xss：XSS攻击   - robot：恶意爬虫   - rfi：远程文件包含   - custom_custom：精准防护   - webshell：网站木马   - custom_whiteblackip：黑白名单拦截   - custom_geoip：地理访问控制拦截   - antitamper：防篡改   - anticrawler：反爬虫    - leakage：网站信息防泄漏   - illegal：非法请求

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
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventItems struct{}"
	}

	return strings.Join([]string{"ListEventItems", string(data)}, " ")
}
