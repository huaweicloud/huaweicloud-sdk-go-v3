package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DimensionResp **参数解释**： 指标维度。
type DimensionResp struct {

	// **参数解释**： 监控维度名称，如ECS的维度为instance_id。各服务资源的维度名称，请参阅具体云服务的文档。您可以直接从[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)页面导航至相应文档。 **取值范围**： 以字母开头，只能包含字母、数字、“_”、“-”。长度[1,32]个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 监控维度取值，例如ECS的ID。 **取值范围**： 长度为[0,256]个字符。
	Value *string `json:"value,omitempty"`
}

func (o DimensionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimensionResp struct{}"
	}

	return strings.Join([]string{"DimensionResp", string(data)}, " ")
}
