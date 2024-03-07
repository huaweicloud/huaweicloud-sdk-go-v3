package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserMetric struct {

	// 用户名称
	Username *string `json:"username,omitempty"`

	// 桌面使用统计信息 * `user_usage` -  用户使用时长(单位:秒)，同一时间登录多台PC的话;相应的时间会累加 * `user_login_count` -  用户登录次数(单位:次) * `user_login_success_count` -  用户登录成功次数(单位:次) * `user_login_fail_count` -  用户登录失败次数(单位:次)
	Metric *[]Metric `json:"metric,omitempty"`
}

func (o UserMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserMetric struct{}"
	}

	return strings.Join([]string{"UserMetric", string(data)}, " ")
}
