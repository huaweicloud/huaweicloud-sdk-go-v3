package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteInstanceDemandRequest struct {

	// **参数解释**： 实例ID。可通过查询实例列表接口ID字段获取 **约束限制**： 不涉及 **取值范围**： 以查询实例列表接口值为准，字符长度32-64。 **默认取值**： 不涉及
	Id string `json:"id"`

	// 是否删除弹性IP
	DeletePublicip *bool `json:"delete_publicip,omitempty"`

	// 是否删除磁盘
	DeleteVolume *bool `json:"delete_volume,omitempty"`
}

func (o DeleteInstanceDemandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceDemandRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceDemandRequest", string(data)}, " ")
}
