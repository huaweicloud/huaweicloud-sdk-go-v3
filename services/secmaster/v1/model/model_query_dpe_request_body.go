package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryDpeRequestBody 查询分类映射请求体
type QueryDpeRequestBody struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 开始时间
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// 结束时间
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// **参数解释：** 偏移量 **约束限制：** 0-10000 **取值范围：** 不涉及 **默认取值：** 0
	Offset int32 `json:"offset"`

	// **参数解释**: 当前页码 **约束限制**: 不涉及
	Limit int32 `json:"limit"`
}

func (o QueryDpeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDpeRequestBody struct{}"
	}

	return strings.Join([]string{"QueryDpeRequestBody", string(data)}, " ")
}
