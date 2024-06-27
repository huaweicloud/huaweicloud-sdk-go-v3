package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmTemplateInfo struct {

	// 创建时间
	CreateTime *string `json:"createTime,omitempty"`

	// 创建者
	CreateUser *string `json:"createUser,omitempty"`

	// UUID
	Id *string `json:"id,omitempty"`

	// 备注
	Remarks *string `json:"remarks,omitempty"`

	// 服务id
	TestServiceId *string `json:"testServiceId,omitempty"`

	// 修改时间
	UpdateTime *string `json:"updateTime,omitempty"`

	// 修改者
	UpdateUser *string `json:"updateUser,omitempty"`

	// 模板名称
	Name *string `json:"name,omitempty"`
}

func (o AlarmTemplateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmTemplateInfo struct{}"
	}

	return strings.Join([]string{"AlarmTemplateInfo", string(data)}, " ")
}
