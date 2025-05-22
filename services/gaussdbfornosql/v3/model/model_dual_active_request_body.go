package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DualActiveRequestBody 参数解释 搭建双活的目标实例信息。 约束限制 不涉及。 取值范围 不涉及。 默认取值 不涉及。
type DualActiveRequestBody struct {

	// 参数解释 搭建双活目标实例所在的region。 约束限制 不涉及。 取值范围 不涉及。 默认取值 不涉及。
	DestinationRegion string `json:"destination_region"`

	// 参数解释 搭建双活目标实例ID。 约束限制 不涉及。 取值范围 不涉及。 默认取值 不涉及。
	DestinationInstanceId string `json:"destination_instance_id"`
}

func (o DualActiveRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DualActiveRequestBody struct{}"
	}

	return strings.Join([]string{"DualActiveRequestBody", string(data)}, " ")
}
