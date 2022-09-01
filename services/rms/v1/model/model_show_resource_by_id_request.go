package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowResourceByIdRequest struct {

	// 云服务名称
	Provider string `json:"provider" xml:"provider"`

	// 资源类型名称
	Type string `json:"type" xml:"type"`

	// 资源ID
	ResourceId string `json:"resource_id" xml:"resource_id"`
}

func (o ShowResourceByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceByIdRequest", string(data)}, " ")
}
