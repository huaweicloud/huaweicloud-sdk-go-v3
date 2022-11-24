package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 项目成员信息
type BindUserRsp struct {

	// 项目成员用户id
	Id *string `json:"id,omitempty"`

	// 项目成员用户名
	Name *string `json:"name,omitempty"`
}

func (o BindUserRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindUserRsp struct{}"
	}

	return strings.Join([]string{"BindUserRsp", string(data)}, " ")
}
