package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesRequest Request Object
type ListResourcesRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 分页参数:每页限定数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数：页数
	Offset *int32 `json:"offset,omitempty"`

	// 资源名称
	ResourceName *string `json:"resourceName,omitempty"`
}

func (o ListResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesRequest", string(data)}, " ")
}
