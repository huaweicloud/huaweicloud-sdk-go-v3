package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NamespaceAllowEmptyResp **参数解释**： 各服务命名空间，请参阅[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)。 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。长度为[0,32]个字符。
type NamespaceAllowEmptyResp struct {
}

func (o NamespaceAllowEmptyResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamespaceAllowEmptyResp struct{}"
	}

	return strings.Join([]string{"NamespaceAllowEmptyResp", string(data)}, " ")
}
