/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type ApplyConfigurationResponseApplyResults struct {
	// 实例ID。
	InstanceId string `json:"instance_id"`
	// 实例名称。
	InstanceName string `json:"instance_name"`
	// 实例是否需要重启。  - “true”需要重启。 - “false”不需要重启。
	RestartRequired bool `json:"restart_required"`
	// 参数模板是否应用成功。  - “true”表示参数模板应用成功。 - “false”表示参数模板应用失败。
	Success bool `json:"success"`
}

func (o ApplyConfigurationResponseApplyResults) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ApplyConfigurationResponseApplyResults", string(data)}, " ")
}
