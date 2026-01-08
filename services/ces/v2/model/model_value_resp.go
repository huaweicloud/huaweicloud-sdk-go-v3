package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValueResp **参数解释**： 告警阈值。具体阈值取值请参见[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)。  **取值范围**： 最小值为-1.7976931348623157e+108，最大值为1.7976931348623157e+108。
type ValueResp struct {
}

func (o ValueResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueResp struct{}"
	}

	return strings.Join([]string{"ValueResp", string(data)}, " ")
}
