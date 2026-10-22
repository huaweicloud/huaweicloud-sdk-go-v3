package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiAppUpdateRequest struct {

	// AI应用名称
	Name *string `json:"name,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`
}

func (o AiAppUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiAppUpdateRequest struct{}"
	}

	return strings.Join([]string{"AiAppUpdateRequest", string(data)}, " ")
}
