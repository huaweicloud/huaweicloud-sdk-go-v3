package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRedisPitrPolicyResponse Response Object
type ShowRedisPitrPolicyResponse struct {

	// 标识Redis实例是否开启指定时间点恢复。 “true”，表示实例开启Redis指定时间点恢复功能。 “false”，表示实例不启用Redis指定时间点恢复功能。
	Enabled *bool `json:"enabled,omitempty"`

	// 数据备份的时间间隔，该数据备份控制redis实例可恢复时间点的间隔，仅在开启时返回。
	Interval *int32 `json:"interval,omitempty"`

	// 指定已生成的备份文件可以保存的天数，仅在开启时返回。
	KeepDays       *int32 `json:"keep_days,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowRedisPitrPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRedisPitrPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowRedisPitrPolicyResponse", string(data)}, " ")
}
