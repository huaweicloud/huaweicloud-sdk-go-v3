package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyRule struct {

	// 源地址网段
	Source *string `json:"source,omitempty"`

	// 目的地址网段
	Destination *[]string `json:"destination,omitempty"`
}

func (o PolicyRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyRule struct{}"
	}

	return strings.Join([]string{"PolicyRule", string(data)}, " ")
}
