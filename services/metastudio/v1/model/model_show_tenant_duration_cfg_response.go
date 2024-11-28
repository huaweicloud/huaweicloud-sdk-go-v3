package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTenantDurationCfgResponse Response Object
type ShowTenantDurationCfgResponse struct {

	// 租户id
	TenantId *string `json:"tenant_id,omitempty"`

	// 基础版最低时长（秒）
	BasicMin *int32 `json:"basic_min,omitempty"`

	// 基础版最高时长（秒）
	BasicMax *int32 `json:"basic_max,omitempty"`

	// 建议时长（秒）
	BasicAdviceValue *int32 `json:"basic_advice_value,omitempty"`

	// 进阶版最低时长（秒）
	MiddleMin *int32 `json:"middle_min,omitempty"`

	// 进阶版最高时长（秒）
	MiddleMax *int32 `json:"middle_max,omitempty"`

	// 建议时长（秒）
	MiddleAdviceValue *int32 `json:"middle_advice_value,omitempty"`

	// 高级版最低时长（秒）
	AdvanceMin *int32 `json:"advance_min,omitempty"`

	// 高级版最高时长（秒）
	AdvanceMax *int32 `json:"advance_max,omitempty"`

	// 建议时长（秒）
	AdvanceAdviceValue *int32 `json:"advance_advice_value,omitempty"`

	// flexus版最低时长（秒）
	FlexusMin *int32 `json:"flexus_min,omitempty"`

	// flexus版最高时长（秒）
	FlexusMax *int32 `json:"flexus_max,omitempty"`

	// flexus建议时长（秒）
	FlexusAdviceValue *int32 `json:"flexus_advice_value,omitempty"`

	// 出门问问最低时长（秒）
	CmwwMin *int32 `json:"cmww_min,omitempty"`

	// 出门问问最高时长（秒）
	CmwwMax *int32 `json:"cmww_max,omitempty"`

	// 出门问问建议时长（秒）
	CmwwAdviceValue *int32 `json:"cmww_advice_value,omitempty"`

	// 逻辑智能最低时长（秒）
	LjznMin *int32 `json:"ljzn_min,omitempty"`

	// 逻辑智能最高时长（秒）
	LjznMax *int32 `json:"ljzn_max,omitempty"`

	// 逻辑智能建议时长（秒）
	LjznAdviceValue *int32 `json:"ljzn_advice_value,omitempty"`
	HttpStatusCode  int    `json:"-"`
}

func (o ShowTenantDurationCfgResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantDurationCfgResponse struct{}"
	}

	return strings.Join([]string{"ShowTenantDurationCfgResponse", string(data)}, " ")
}
