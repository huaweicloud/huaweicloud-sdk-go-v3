package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkOrderDetailVo struct {

	// 业务受理ID
	Id *int64 `json:"id,omitempty"`

	// SIM卡类型:3.实体卡
	SimType *int32 `json:"sim_type,omitempty"`

	// 业务受理明细状态：1成功、2处理中、3失败
	Status *int32 `json:"status,omitempty"`

	// 容器ID
	Cid *string `json:"cid,omitempty"`

	// SIM卡标识
	SimCardId *int64 `json:"sim_card_id,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 完成时间
	FinishTime *sdktime.SdkTime `json:"finish_time,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o WorkOrderDetailVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkOrderDetailVo struct{}"
	}

	return strings.Join([]string{"WorkOrderDetailVo", string(data)}, " ")
}
