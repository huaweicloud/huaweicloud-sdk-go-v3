package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowResourceByIdRequest struct {
	// 云服务英文简写

	Provider string `json:"provider"`
	// 云服务资源类型名称

	Type string `json:"type"`
	// 资源ID

	ResourceId string `json:"resource_id"`
}

func (o ShowResourceByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceByIdRequest", string(data)}, " ")
}
