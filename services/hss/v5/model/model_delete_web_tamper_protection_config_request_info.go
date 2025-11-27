package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteWebTamperProtectionConfigRequestInfo struct {

	// **参数解释**: 网页防篡改的类型 **约束限制**: 不涉及 **取值范围**: - container_wtp：容器网页防篡改  **默认取值**: 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**: 防护配置id列表 **约束限制**: 不涉及 **取值范围**: 最少0条，最多100条 **默认取值**: 不涉及
	ProtectionConfigIdList *[]string `json:"protection_config_id_list,omitempty"`
}

func (o DeleteWebTamperProtectionConfigRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWebTamperProtectionConfigRequestInfo struct{}"
	}

	return strings.Join([]string{"DeleteWebTamperProtectionConfigRequestInfo", string(data)}, " ")
}
