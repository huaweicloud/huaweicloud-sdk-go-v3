package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Group 群组信息
type Group struct {

	// 群组ID
	Id *string `json:"id,omitempty"`

	// 群组名称
	Name *string `json:"name,omitempty"`
}

func (o Group) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Group struct{}"
	}

	return strings.Join([]string{"Group", string(data)}, " ")
}
