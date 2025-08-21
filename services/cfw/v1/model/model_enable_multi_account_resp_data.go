package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableMultiAccountRespData **参数解释**： 响应信息 **取值范围**： 不涉及
type EnableMultiAccountRespData struct {

	// **参数解释**： 防火墙ID **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 防火墙名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 云防火墙可信服务状态 **取值范围**： 1 已开启
	TrustServiceStatus *int32 `json:"trust_service_status,omitempty"`
}

func (o EnableMultiAccountRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableMultiAccountRespData struct{}"
	}

	return strings.Join([]string{"EnableMultiAccountRespData", string(data)}, " ")
}
