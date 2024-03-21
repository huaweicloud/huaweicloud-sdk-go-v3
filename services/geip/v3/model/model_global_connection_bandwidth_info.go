package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GlobalConnectionBandwidthInfo struct {

	// 骨干带宽的ID
	GcbId *string `json:"gcb_id,omitempty"`

	// 骨干带宽的大小
	Size *int32 `json:"size,omitempty"`

	// 骨干带宽类型（城域、区域和大区）
	GcbType *string `json:"gcb_type,omitempty"`

	// - 功能说明：骨干带宽状态 - 取值范围：NORMAL 正常、FREEZED 冻结
	AdminState *string `json:"admin_state,omitempty"`

	// - 功能说明：网络服务等级 - 取值范围：Pt - 铂金，Au - 金牌，Ag - 银牌，Cu - 铜牌
	SlaLevel *string `json:"sla_level,omitempty"`

	// 线路质量金银铜对应的DSCP值
	Dscp *int32 `json:"dscp,omitempty"`
}

func (o GlobalConnectionBandwidthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalConnectionBandwidthInfo struct{}"
	}

	return strings.Join([]string{"GlobalConnectionBandwidthInfo", string(data)}, " ")
}
