package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmHistoryItemV2MetricDimensions **参数解释**： 资源信息。 **取值范围**： 不涉及。
type AlarmHistoryItemV2MetricDimensions struct {

	// **参数解释**： 资源维度名称，如：弹性云服务器，则维度为instance_id。各服务资源的维度名称，请参阅具体云服务的文档。您可以直接从[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)页面导航至相应文档。。 **取值范围**： 以字母开头，只能包含字母、数字、“_”、“-”。长度为[1,32]个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 资源维度值，为资源的实例ID，如：4270ff17-aba3-4138-89fa-820594c39755。 **取值范围**： 字符串长度在 1 到 256 之间。
	Value *string `json:"value,omitempty"`
}

func (o AlarmHistoryItemV2MetricDimensions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmHistoryItemV2MetricDimensions struct{}"
	}

	return strings.Join([]string{"AlarmHistoryItemV2MetricDimensions", string(data)}, " ")
}
