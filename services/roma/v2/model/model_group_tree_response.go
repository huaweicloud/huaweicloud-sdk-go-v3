package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupTreeResponse struct {

	// 分组id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 分组名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 父分组id
	ParentId *int32 `json:"parent_id,omitempty" xml:"parent_id"`

	// 子分组
	Children *[]GroupTreeResponse `json:"children,omitempty" xml:"children"`

	// 应用id
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 权限
	Permissions *[]string `json:"permissions,omitempty" xml:"permissions"`
}

func (o GroupTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupTreeResponse struct{}"
	}

	return strings.Join([]string{"GroupTreeResponse", string(data)}, " ")
}
