package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ManageableGroupDto struct {

	// 代码组全名
	FullName *string `json:"full_name,omitempty"`

	// 代码组id
	Id *int32 `json:"id,omitempty"`

	// 代码组名
	Name *string `json:"name,omitempty"`
}

func (o ManageableGroupDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ManageableGroupDto struct{}"
	}

	return strings.Join([]string{"ManageableGroupDto", string(data)}, " ")
}
