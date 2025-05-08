package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupCreateReq struct {

	// 消费组名称
	GroupName *string `json:"group_name,omitempty"`

	// 消费组描述
	GroupDesc *string `json:"group_desc,omitempty"`
}

func (o GroupCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupCreateReq struct{}"
	}

	return strings.Join([]string{"GroupCreateReq", string(data)}, " ")
}
