package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociatedResourceListResponse struct {

	// 资源ID
	Id *string `json:"id,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// eip信息
	Eip *string `json:"eip,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`
}

func (o AssociatedResourceListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatedResourceListResponse struct{}"
	}

	return strings.Join([]string{"AssociatedResourceListResponse", string(data)}, " ")
}
