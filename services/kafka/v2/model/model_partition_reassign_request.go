package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PartitionReassignRequest struct {

	// 分区平衡分配方案。
	Reassignments []PartitionReassignEntity `json:"reassignments"`

	// 分区平衡门限值。
	Throttle *int32 `json:"throttle,omitempty"`

	// **参数解释**： 是否作为定时任务执行。若非定时执行，is_schedule和execute_at字段可为空。若为定时执行，is_schedule为true，execute_at字段非空。 **约束限制**： [不涉及。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm,dt,ax,sbc,hk_sbc,srg,fcs,cmcc)[华为云Stack不支持此参数。](tag:hcs)
	IsSchedule *bool `json:"is_schedule,omitempty"`

	// **参数解释**： 定时时间，格式为Unix时间戳，单位为毫秒 **约束限制**： [不涉及。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm,dt,ax,sbc,hk_sbc,srg,fcs,cmcc)[华为云Stack不支持此参数。](tag:hcs)
	ExecuteAt *int64 `json:"execute_at,omitempty"`

	// **参数解释**： 设为true表示执行时间预估任务，false为执行分区平衡任务。 **约束限制**： [不涉及。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm,dt,ax,sbc,hk_sbc,srg,fcs,cmcc)[华为云Stack不支持此参数。](tag:hcs)
	TimeEstimate *bool `json:"time_estimate,omitempty"`
}

func (o PartitionReassignRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PartitionReassignRequest struct{}"
	}

	return strings.Join([]string{"PartitionReassignRequest", string(data)}, " ")
}
