package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchTokenResponseAk 授权的ak数据
type SwitchTokenResponseAk struct {

	// ak名称
	Name *string `json:"name,omitempty"`

	// ak id
	Id *string `json:"id,omitempty"`

	// ak状态
	Status *string `json:"status,omitempty"`
}

func (o SwitchTokenResponseAk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchTokenResponseAk struct{}"
	}

	return strings.Join([]string{"SwitchTokenResponseAk", string(data)}, " ")
}
