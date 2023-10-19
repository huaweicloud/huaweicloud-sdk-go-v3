package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ParamGroupHistoryResponse struct {

	// 参数名称。
	ParameterName *string `json:"parameter_name,omitempty"`

	// 修改前参数值。
	OldValue *string `json:"old_value,omitempty"`

	// 修改后参数值。
	NewValue *string `json:"new_value,omitempty"`

	// 更新结果。
	UpdateResult *string `json:"update_result,omitempty"`

	// 是否应用。 - true：是。 - false：否。
	IsApplied *bool `json:"is_applied,omitempty"`

	// 修改时间。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	Updated *string `json:"updated,omitempty"`

	// 应用时间。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	Applied *string `json:"applied,omitempty"`
}

func (o ParamGroupHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParamGroupHistoryResponse struct{}"
	}

	return strings.Join([]string{"ParamGroupHistoryResponse", string(data)}, " ")
}
