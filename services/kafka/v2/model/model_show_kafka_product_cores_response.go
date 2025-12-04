package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKafkaProductCoresResponse Response Object
type ShowKafkaProductCoresResponse struct {

	// **参数解释**： 核数。 **取值范围**： 不涉及。
	CoreNum *int32 `json:"core_num,omitempty"`

	// **参数解释**： 需要扩容的存储空间。 **取值范围**： 不涉及。
	TotalExtendStorageSpace *int32 `json:"total_extend_storage_space,omitempty"`
	HttpStatusCode          int    `json:"-"`
}

func (o ShowKafkaProductCoresResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaProductCoresResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaProductCoresResponse", string(data)}, " ")
}
