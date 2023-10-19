package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyDelta struct {
	Policy *Policy `json:"policy,omitempty"`

	// 变更类型
	ChangeType *int32 `json:"change_type,omitempty"`
}

func (o PolicyDelta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyDelta struct{}"
	}

	return strings.Join([]string{"PolicyDelta", string(data)}, " ")
}
