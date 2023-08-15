package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ErrorCount struct {

	// 为空。
	Empty *bool `json:"empty,omitempty"`
}

func (o ErrorCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorCount struct{}"
	}

	return strings.Join([]string{"ErrorCount", string(data)}, " ")
}
