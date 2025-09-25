package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CancelHostsQuotaRequestInfo struct {

	// **参数解释**: 资源ID列表 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及
	ResourceIdList []string `json:"resource_id_list"`
}

func (o CancelHostsQuotaRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelHostsQuotaRequestInfo struct{}"
	}

	return strings.Join([]string{"CancelHostsQuotaRequestInfo", string(data)}, " ")
}
