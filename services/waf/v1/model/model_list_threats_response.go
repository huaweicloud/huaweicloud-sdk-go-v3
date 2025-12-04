package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListThreatsResponse Response Object
type ListThreatsResponse struct {

	// 时间区间内xss攻击数量
	Xss *int32 `json:"xss,omitempty"`

	// 时间区间内sqli攻击数量
	Sqli *int32 `json:"sqli,omitempty"`

	// 时间区间内cmdi攻击数量
	Cmdi *int32 `json:"cmdi,omitempty"`

	// 时间区间内lfi攻击数量
	Lfi *int32 `json:"lfi,omitempty"`

	// 时间区间内rfi攻击数量
	Rfi *int32 `json:"rfi,omitempty"`

	// 时间区间内webshell攻击数量
	Webshell *int32 `json:"webshell,omitempty"`

	// 时间区间内恶意爬虫数量
	Robot *int32 `json:"robot,omitempty"`

	// 时间区间内cc攻击数量
	Cc *int32 `json:"cc,omitempty"`

	// 时间区间内对自定义规则攻击数量
	Custom *int32 `json:"custom,omitempty"`

	// 时间区间内对黑白名单规则攻击数量
	Whiteblackip *int32 `json:"whiteblackip,omitempty"`

	// 时间区间内反泄漏数量
	Antileakage *int32 `json:"antileakage,omitempty"`

	// 时间区间内防篡改数量
	Antitamper     *int32 `json:"antitamper,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListThreatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListThreatsResponse struct{}"
	}

	return strings.Join([]string{"ListThreatsResponse", string(data)}, " ")
}
