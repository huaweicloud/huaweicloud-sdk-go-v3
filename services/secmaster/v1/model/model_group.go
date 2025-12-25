package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Group Group具体内容
type Group struct {

	// UUID
	GroupId *string `json:"group_id,omitempty"`

	// 所属租户名称
	Name *string `json:"name,omitempty"`
}

func (o Group) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Group struct{}"
	}

	return strings.Join([]string{"Group", string(data)}, " ")
}
