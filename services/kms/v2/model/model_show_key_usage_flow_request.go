package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKeyUsageFlowRequest Request Object
type ShowKeyUsageFlowRequest struct {

	// **参数解释：** 密钥空间ID **约束限制：** 满足正则表达式^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$ **取值范围：** 不涉及 **默认取值：** 不涉及
	KeyspaceId string `json:"keyspace_id"`

	// **参数解释：** 密钥ID **约束限制：** 满足正则表达式^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$ **取值范围：** 不涉及 **默认取值：** 不涉及
	KeyId string `json:"key_id"`
}

func (o ShowKeyUsageFlowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKeyUsageFlowRequest struct{}"
	}

	return strings.Join([]string{"ShowKeyUsageFlowRequest", string(data)}, " ")
}
