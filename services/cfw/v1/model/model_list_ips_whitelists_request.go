package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpsWhitelistsRequest Request Object
type ListIpsWhitelistsRequest struct {

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 白名单源地址 **约束限制**： 不涉及  **取值范围**： 不涉及  **默认取值**： 不涉及
	SourceAddress *string `json:"source_address,omitempty"`

	// **参数解释**：  目的地址  **约束限制**：  不涉及  **取值范围**：  不涉及  **默认取值**：  不涉及
	DestAddress *string `json:"dest_address,omitempty"`

	// **参数解释**：  白名单名称  **约束限制**：  不涉及  **取值范围**： 不涉及  **默认取值**：  不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：  每页大小  **约束限制**：  不涉及  **取值范围**：  不涉及  **默认取值**：  不涉及
	Limit int32 `json:"limit"`

	// **参数解释**：  偏移量  **约束限制**：  不涉及  **取值范围**：  不涉及  **默认取值**：  不涉及
	Offset int32 `json:"offset"`
}

func (o ListIpsWhitelistsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpsWhitelistsRequest struct{}"
	}

	return strings.Join([]string{"ListIpsWhitelistsRequest", string(data)}, " ")
}
