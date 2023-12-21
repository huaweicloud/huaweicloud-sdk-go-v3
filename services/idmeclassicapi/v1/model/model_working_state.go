package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkingState struct {

	// 别名。
	Alias *string `json:"alias,omitempty"`

	// 中文名称。
	CnName *string `json:"cnName,omitempty"`

	// 编码。
	Code *string `json:"code,omitempty"`

	// 英文名称。
	EnName *string `json:"enName,omitempty"`
}

func (o WorkingState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkingState struct{}"
	}

	return strings.Join([]string{"WorkingState", string(data)}, " ")
}
