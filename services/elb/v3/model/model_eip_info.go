package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 弹性ip，同publicips
type EipInfo struct {
	// eip_id

	EipId *string `json:"eip_id,omitempty"`
	// eip_address

	EipAddress *string `json:"eip_address,omitempty"`
}

func (o EipInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipInfo struct{}"
	}

	return strings.Join([]string{"EipInfo", string(data)}, " ")
}
