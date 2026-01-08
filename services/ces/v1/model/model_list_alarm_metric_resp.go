package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmMetricResp **参数解释**： 创建规则中的监控指标信息。
type ListAlarmMetricResp struct {

	// **参数解释**： 服务指标命名空间。如：弹性云服务器的命名空间为SYS.ECS，文档数据库的命名空间为SYS.DDS，各服务命名空间请参阅[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)。 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_，字符总长度最短为3，最大为32。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**： 指标维度，目前最大为4个维度。
	Dimensions *[]DimensionResp `json:"dimensions,omitempty"`

	// **参数解释**： 资源的监控指标名称。如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率；各服务资源的指标名称，请参阅具体云服务的文档。您可以直接从[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)页面导航至相应文档。 **取值范围**： 必须以字母开头，只能包含0-9/a-z/A-Z/_，长度为[1,96]个字符。
	MetricName *string `json:"metric_name,omitempty"`

	// **参数解释**： 创建告警规则时选择的资源分组ID，如：rg1603786526428bWbVmk4rP **取值范围**： 只能包含字母、数字、“-”、“_”，字符长度为[1,64]
	ResourceGroupId *string `json:"resource_group_id,omitempty"`

	// **参数解释**： 创建告警规则时选择的资源分组名称，如：Resource-Group-ECS-01 **取值范围**： 字符长度为[0,128]
	ResourceGroupName *string `json:"resource_group_name,omitempty"`
}

func (o ListAlarmMetricResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmMetricResp struct{}"
	}

	return strings.Join([]string{"ListAlarmMetricResp", string(data)}, " ")
}
