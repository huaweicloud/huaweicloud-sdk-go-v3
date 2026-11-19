package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowKeyUsageFlowResponseBodyNodes struct {

	// **参数解释：** 密钥节点 **取值范围：** 不涉及
	Key *interface{} `json:"key,omitempty"`
}

func (o ShowKeyUsageFlowResponseBodyNodes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKeyUsageFlowResponseBodyNodes struct{}"
	}

	return strings.Join([]string{"ShowKeyUsageFlowResponseBodyNodes", string(data)}, " ")
}
