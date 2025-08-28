package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PoolHealth struct {

	// **参数解释**：当健康检查在线的member个数小于该个数，判定pool的state为不健康。  **取值范围**： - 0：默认值，不生效。 - 1：全下线转发生效。  [不支持该字段，请勿使用。](tag:hws_eu,hws_eu_wb,hws_test,fcs,dt,hcso_dt,ctc,cmcc,tm,sbc,hk_sbc,hk_tm,hk_vdf,srg,g42,hk_g42)
	MinimumHealthyMemberCount *int32 `json:"minimum_healthy_member_count,omitempty"`
}

func (o PoolHealth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolHealth struct{}"
	}

	return strings.Join([]string{"PoolHealth", string(data)}, " ")
}
