package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCustomizeClusterTagsByProjectIdRequest Request Object
type ShowCustomizeClusterTagsByProjectIdRequest struct {

	// 资源类型，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ResourceType string `json:"resource_type"`
}

func (o ShowCustomizeClusterTagsByProjectIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomizeClusterTagsByProjectIdRequest struct{}"
	}

	return strings.Join([]string{"ShowCustomizeClusterTagsByProjectIdRequest", string(data)}, " ")
}
