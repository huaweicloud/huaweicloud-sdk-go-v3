package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationHistoryRsp struct {

	// 参数名称。
	ParameterName string `json:"parameter_name"`

	// 参数旧值
	OldValue string `json:"old_value"`

	// 参数新值
	NewValue string `json:"new_value"`

	// 更新结果
	UpdateResult string `json:"update_result"`

	// - true:已生效 - false:未生效
	Applied bool `json:"applied"`

	// 更新时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。  [其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。](tag:hc)  [其中，T指某个时间的开始；Z指时区偏移量。](tag:hk)
	UpdatedAt string `json:"updated_at"`

	// 生效时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。  [其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。](tag:hc)  [其中，T指某个时间的开始；Z指时区偏移量。](tag:hk)
	AppliedAt string `json:"applied_at"`
}

func (o ConfigurationHistoryRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationHistoryRsp struct{}"
	}

	return strings.Join([]string{"ConfigurationHistoryRsp", string(data)}, " ")
}
