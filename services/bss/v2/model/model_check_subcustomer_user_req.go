package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckSubcustomerUserReq struct {

	// 该字段内容可填为：“email”、“mobile”或“name”。
	SearchType string `json:"search_type"`

	// 手机号、邮箱或登录名称。 手机号需符合正则表达式 ^\\d{4}-\\d+$；包括国家码，以00开头，格式：00XX-XXXXXXXX。目前手机号仅支持以86开头的国家码。邮箱需为含有@的正确格式的完整邮箱地址。name：符合正则表达式^([a-zA-Z-]([a-zA-Z0-9_-]){4,31})$，长度5-32；不能以“op_”或“shadow_”开头且不能全为数字，且只能以字母（不区分大小写）或者-开头。
	SearchValue string `json:"search_value"`
}

func (o CheckSubcustomerUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckSubcustomerUserReq struct{}"
	}

	return strings.Join([]string{"CheckSubcustomerUserReq", string(data)}, " ")
}
