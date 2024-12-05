package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDisasterRecoverySettingsRequest Request Object
type ShowDisasterRecoverySettingsRequest struct {

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 索引位置偏移量，表示从指定offset条数据后查询对应的实例信息。 取值大于或等于0。不传该参数时，查询偏移量默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询实例个数上限值。 取值范围：1~50。不传该参数时，默认查询前50条实例信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowDisasterRecoverySettingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDisasterRecoverySettingsRequest struct{}"
	}

	return strings.Join([]string{"ShowDisasterRecoverySettingsRequest", string(data)}, " ")
}
