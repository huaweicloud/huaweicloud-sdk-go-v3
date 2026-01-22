package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRocketMqProductCoresResponse Response Object
type ShowRocketMqProductCoresResponse struct {

	// **参数解释**： 核数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CoreNum *int32 `json:"core_num,omitempty"`

	// **参数解释**： 预估存储空间，当填写的broker_num小于等于当前实例真实值时，显示为当前实例的存储空间。如果填写的broker_num大于当前实例真实值时，显示为所填写broker_num时实例的预估存储空间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TotalExtendStorageSpace *int32 `json:"total_extend_storage_space,omitempty"`
	HttpStatusCode          int    `json:"-"`
}

func (o ShowRocketMqProductCoresResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRocketMqProductCoresResponse struct{}"
	}

	return strings.Join([]string{"ShowRocketMqProductCoresResponse", string(data)}, " ")
}
