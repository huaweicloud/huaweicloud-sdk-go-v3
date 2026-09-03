package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotSetChargeModeInstanceResponse Response Object
type ListNotSetChargeModeInstanceResponse struct {

	// 实例列表
	InstanceList *[]InstanceSimpleDto `json:"instance_list,omitempty"`

	// 付费状态。取值范围：0（免费实例）、1（付费实例）
	QuotaStatus *int32 `json:"quota_status,omitempty"`

	// 开通配额总数
	QuotaNum *int32 `json:"quota_num,omitempty"`

	// 已使用配额数量
	UsedNum        *int32 `json:"used_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListNotSetChargeModeInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotSetChargeModeInstanceResponse struct{}"
	}

	return strings.Join([]string{"ListNotSetChargeModeInstanceResponse", string(data)}, " ")
}
