package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpGroupRelatedListenersRequest Request Object
type ShowIpGroupRelatedListenersRequest struct {

	// **参数解释**：IP地址组ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	IpgroupId string `json:"ipgroup_id"`
}

func (o ShowIpGroupRelatedListenersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpGroupRelatedListenersRequest struct{}"
	}

	return strings.Join([]string{"ShowIpGroupRelatedListenersRequest", string(data)}, " ")
}
