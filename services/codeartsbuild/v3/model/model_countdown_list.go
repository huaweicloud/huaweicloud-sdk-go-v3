package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountdownList 构建成功率
type CountdownList struct {

	// 服务类型
	ServiceType *string `json:"service_type,omitempty"`

	// 资源编号
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源规格
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	Tips *CountdownListTips `json:"tips,omitempty"`
}

func (o CountdownList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountdownList struct{}"
	}

	return strings.Join([]string{"CountdownList", string(data)}, " ")
}
