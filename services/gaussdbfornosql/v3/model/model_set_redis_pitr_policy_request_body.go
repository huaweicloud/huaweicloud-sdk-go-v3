package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetRedisPitrPolicyRequestBody struct {

	// 标识Redis实例是否开启指定时间点恢复。 “true”，表示实例开启Redis指定时间点恢复功能。 “false”，表示实例不启用Redis指定时间点恢复功能。
	Enabled bool `json:"enabled"`

	// 数据备份的时间间隔，该数据备份控制redis实例可恢复时间点的间隔，默认值为 20分钟。 取值范围：5～120  单位：分钟  例如，当interval 为20min时，可恢复时间点的间隔为20min，其interval约小，对性能影响越大，存储空间膨胀约明显。
	Interval *int32 `json:"interval,omitempty"`

	// 指定已生成的备份文件可以保存的天数，默认值为 1 天。 取值范围：1～7  单位：天
	KeepDays *int32 `json:"keep_days,omitempty"`
}

func (o SetRedisPitrPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRedisPitrPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"SetRedisPitrPolicyRequestBody", string(data)}, " ")
}
