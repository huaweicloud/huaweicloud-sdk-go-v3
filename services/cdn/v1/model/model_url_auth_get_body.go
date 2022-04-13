package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// URL鉴权查询响应体
type UrlAuthGetBody struct {
	// A/B/C类防盗链开关（\"off\"/\"on\"）。

	Status string `json:"status"`
	// 鉴权方式 type_a：鉴权方式A type_b：鉴权方式B type_c1：鉴权方式C1 type_c2：鉴权方式C2

	Type *string `json:"type,omitempty"`
	// 时间格式 dec：十进制 hex：十六进制 鉴权方式A：只支持十进制 鉴权方式B：只支持十进制 鉴权方式C1：只支持十六进制鉴权方式 鉴权方式C2：支持十进制/十六进制

	TimeFormat *string `json:"time_format,omitempty"`
	// 过期时间：范围：0-31536000单位为秒。

	ExpireTime *int32 `json:"expire_time,omitempty"`
}

func (o UrlAuthGetBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlAuthGetBody struct{}"
	}

	return strings.Join([]string{"UrlAuthGetBody", string(data)}, " ")
}
