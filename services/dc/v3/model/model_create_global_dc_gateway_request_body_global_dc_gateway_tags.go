package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateGlobalDcGatewayRequestBodyGlobalDcGatewayTags struct {

	// 键
	Key string `json:"key"`

	// 值
	Value *string `json:"value,omitempty"`
}

func (o CreateGlobalDcGatewayRequestBodyGlobalDcGatewayTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalDcGatewayRequestBodyGlobalDcGatewayTags struct{}"
	}

	return strings.Join([]string{"CreateGlobalDcGatewayRequestBodyGlobalDcGatewayTags", string(data)}, " ")
}
