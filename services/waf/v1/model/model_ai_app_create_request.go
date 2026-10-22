package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiAppCreateRequest struct {

	// AI应用名称
	Name string `json:"name"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 要与应用关联的策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// API签名
	ApiSign *string `json:"api_sign,omitempty"`
}

func (o AiAppCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiAppCreateRequest struct{}"
	}

	return strings.Join([]string{"AiAppCreateRequest", string(data)}, " ")
}
