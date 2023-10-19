package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessServiceInfo 接入服务授权信息
type AccessServiceInfo struct {

	// 接入服务名称。最大长度为20个字符。
	Name string `json:"name"`

	// 接入服务描述。最大长度为500个字符。
	Description *string `json:"description,omitempty"`

	// 是否授权。值为true或false。
	Grant *bool `json:"grant,omitempty"`
}

func (o AccessServiceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessServiceInfo struct{}"
	}

	return strings.Join([]string{"AccessServiceInfo", string(data)}, " ")
}
