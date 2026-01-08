package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductMetric struct {

	// **参数解释**： 按云产品维度屏蔽时的资源维度信息。    **约束限制**： 不涉及。 **取值范围**： 长度为[0,128]个字符，有多个时用\",\"连接。    **默认取值**： 不涉及。
	DimensionName string `json:"dimension_name"`

	// **参数解释**： 资源的监控指标名称，各服务资源的指标名称，请参阅具体云服务的文档。您可以直接从[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)页面导航至相应文档。 **约束限制**： 不涉及。 **取值范围**： 必须以字母开头，只能包含0-9/a-z/A-Z/_/-。字符长度最短为1，最大为96。如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率。         **默认取值**： 不涉及。
	MetricName string `json:"metric_name"`
}

func (o ProductMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductMetric struct{}"
	}

	return strings.Join([]string{"ProductMetric", string(data)}, " ")
}
