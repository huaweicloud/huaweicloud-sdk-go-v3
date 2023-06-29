package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Group IAM用户组信息
type Group struct {

	// 用户组id
	Id *string `json:"id,omitempty"`

	// 用户组名
	Name *string `json:"name,omitempty"`
}

func (o Group) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Group struct{}"
	}

	return strings.Join([]string{"Group", string(data)}, " ")
}
