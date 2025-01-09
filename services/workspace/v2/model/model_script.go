package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Script 脚本信息。
type Script struct {

	// 脚本id。
	Id *string `json:"id,omitempty"`

	// 脚本名称。
	Name *string `json:"name,omitempty"`
}

func (o Script) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Script struct{}"
	}

	return strings.Join([]string{"Script", string(data)}, " ")
}
