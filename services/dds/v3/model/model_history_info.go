package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HistoryInfo struct {

	// 参数名称
	ParameterName string `json:"parameter_name"`

	// 修改前的值。
	OldValue string `json:"old_value"`

	// 修改后的值。
	NewValue string `json:"new_value"`

	// 修改时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	UpdatedAt string `json:"updated_at"`
}

func (o HistoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryInfo struct{}"
	}

	return strings.Join([]string{"HistoryInfo", string(data)}, " ")
}
