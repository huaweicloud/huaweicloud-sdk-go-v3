package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricNameSchema **参数解释**： 指标ID，例如弹性云服务器的监控指标CPU使用率，对应的metric_name为cpu_util。各服务资源的指标名称，请参阅具体云服务的文档。您可以直接从[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)页面导航至相应文档。 **约束限制**： 不涉及。 **取值范围**： 必须以字母开头，只能包含0-9/a-z/A-Z/_/-；如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率。字符串长度为[1,96]。 **默认取值**： 不涉及。
type MetricNameSchema struct {
}

func (o MetricNameSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricNameSchema struct{}"
	}

	return strings.Join([]string{"MetricNameSchema", string(data)}, " ")
}
