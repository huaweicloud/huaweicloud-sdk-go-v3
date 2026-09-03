package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlertTemplate struct {

	// 告警模板id
	Id *string `json:"id,omitempty"`

	// 告警模板名称
	Name *string `json:"name,omitempty"`
}

func (o AlertTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertTemplate struct{}"
	}

	return strings.Join([]string{"AlertTemplate", string(data)}, " ")
}
