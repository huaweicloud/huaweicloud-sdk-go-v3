package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CodeSource struct {

	// 流水线源类型
	Type *string `json:"type,omitempty"`

	Params *CodeSourceParams `json:"params,omitempty"`
}

func (o CodeSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeSource struct{}"
	}

	return strings.Join([]string{"CodeSource", string(data)}, " ")
}
