package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupTreeResponse struct {
	// 分组id

	Id *int32 `json:"id,omitempty"`
	// 分组名称

	Name *string `json:"name,omitempty"`
	// 父分组id

	ParentId *int32 `json:"parent_id,omitempty"`
	// 子分组

	Children *[]GroupTreeResponse `json:"children,omitempty"`
	// 应用id

	AppId *string `json:"app_id,omitempty"`
	// 权限

	Permissions *[]string `json:"permissions,omitempty"`
}

func (o GroupTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupTreeResponse struct{}"
	}

	return strings.Join([]string{"GroupTreeResponse", string(data)}, " ")
}
