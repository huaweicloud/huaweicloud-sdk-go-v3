package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// vpc信息。
type VpcConfig struct {

	// vpc名称。
	VpcName *string `json:"vpc_name,omitempty"`

	// vpc ID。
	VpcId *string `json:"vpc_id,omitempty"`
}

func (o VpcConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcConfig struct{}"
	}

	return strings.Join([]string{"VpcConfig", string(data)}, " ")
}
