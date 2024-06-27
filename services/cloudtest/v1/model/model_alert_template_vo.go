package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlertTemplateVo struct {

	// 告警级别列表
	AlertLevels *[]AlertLevel `json:"alertLevels,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建人
	CreateUser *string `json:"create_user,omitempty"`

	// 唯一ID，主键
	Id *string `json:"id,omitempty"`

	// 告警模板名称
	Name *string `json:"name,omitempty"`

	// 备注
	Remarks *string `json:"remarks,omitempty"`

	// 服务ID
	TestServiceId *string `json:"test_service_id,omitempty"`

	// 创建时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 更新人
	UpdateUser *string `json:"update_user,omitempty"`
}

func (o AlertTemplateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertTemplateVo struct{}"
	}

	return strings.Join([]string{"AlertTemplateVo", string(data)}, " ")
}
