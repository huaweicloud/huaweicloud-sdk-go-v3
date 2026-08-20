package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EndpointCreatorInfo struct {

	// 用户id
	UserId *string `json:"user_id,omitempty"`

	// 用户名称
	Username *string `json:"username,omitempty"`
}

func (o EndpointCreatorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointCreatorInfo struct{}"
	}

	return strings.Join([]string{"EndpointCreatorInfo", string(data)}, " ")
}
