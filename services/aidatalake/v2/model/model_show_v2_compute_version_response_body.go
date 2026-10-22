package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowV2ComputeVersionResponseBody **参数解释**：计算资源版本详情。
type ShowV2ComputeVersionResponseBody struct {

	// **参数解释**：计算资源记录ID。 **取值范围**：长度为1~36的英文字符、数字和中划线的组合。
	Id *string `json:"id,omitempty"`

	// **参数解释**：计算资源ID。 **取值范围**：长度为1~36的英文字母、数字和中划线的组合。
	ComputeId *string `json:"compute_id,omitempty"`

	// **参数解释**：计算资源名称。 **取值范围**：支持小写字母、数字、中划线，必须字母开头、字母或数字结尾，长度为1~63个字符。
	Name *string `json:"name,omitempty"`

	Capacity *CapacityVersion `json:"capacity,omitempty"`

	// **参数解释**：资源创建时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释**：资源更新时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	Creator *Creator `json:"creator,omitempty"`

	Error *Error `json:"error,omitempty"`

	Operation *Operation `json:"operation,omitempty"`
}

func (o ShowV2ComputeVersionResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowV2ComputeVersionResponseBody struct{}"
	}

	return strings.Join([]string{"ShowV2ComputeVersionResponseBody", string(data)}, " ")
}
