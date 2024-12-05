package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetDisasterRecoverySettingsResponse Response Object
type SetDisasterRecoverySettingsResponse struct {

	// 设置容灾切换故障节点比例成功的实例列表。
	SuccessedInstanceIds *[]string `json:"successed_instance_ids,omitempty"`
	HttpStatusCode       int       `json:"-"`
}

func (o SetDisasterRecoverySettingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDisasterRecoverySettingsResponse struct{}"
	}

	return strings.Join([]string{"SetDisasterRecoverySettingsResponse", string(data)}, " ")
}
