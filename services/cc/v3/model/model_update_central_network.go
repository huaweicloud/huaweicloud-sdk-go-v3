package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCentralNetwork 更新中心网络的详细信息。
type UpdateCentralNetwork struct {

	// 实例名字。
	Name *string `json:"name,omitempty"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例标签。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o UpdateCentralNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCentralNetwork struct{}"
	}

	return strings.Join([]string{"UpdateCentralNetwork", string(data)}, " ")
}
