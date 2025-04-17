package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Value 告警阈值。单一阈值时value和alarm_level配对使用，当hierarchical_value和value同时使用时以hierarchical_value为准。取值范围[0, Number.MAX_VALUE]，Number.MAX_VALUE值为1.7976931348623157e+108。具体阈值取值请参见附录中各服务监控指标中取值范围，如支持监控的服务列表中ECS的CPU使用率cpu_util取值范围可配置80。 [具体阈值取值请参见附录中各服务监控指标中取值范围，如[支持监控的服务列表](ces_03_0059.xml)中ECS的CPU使用率cpu_util取值范围可配置80。](tag: dt,g42,dt_test,hk_g42,hk_sbc,hws,hws_hk,ocb,sbc,tm)
type Value struct {
}

func (o Value) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Value struct{}"
	}

	return strings.Join([]string{"Value", string(data)}, " ")
}
