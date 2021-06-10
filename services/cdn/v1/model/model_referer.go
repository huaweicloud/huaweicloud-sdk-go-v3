package model

import (
	"encoding/json"

	"strings"
)

type Referer struct {
	// Referer类型。取值：0代表不设置Referer过滤；1代表黑名单；2代表白名单。

	RefererType int32 `json:"referer_type"`
	// 分号隔开的域名列表。

	RefererList *string `json:"referer_list,omitempty"`
	// 是否包含空Referer。如果是黑名单并开启该选项，则表示无referer不允许访问。如果是白名单并开启该选项，则表示无referer允许访问。

	IncludeEmpty *bool `json:"include_empty,omitempty"`
}

func (o Referer) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Referer struct{}"
	}

	return strings.Join([]string{"Referer", string(data)}, " ")
}
