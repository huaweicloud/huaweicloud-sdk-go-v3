package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDevServerPublicIpRequest Request Object
type ListDevServerPublicIpRequest struct {

	// **参数解释**：Lite Server ID。 **约束限制**：必填。 **取值范围**：1 - 64字符。 **默认取值**：不涉及。
	Id string `json:"id"`
}

func (o ListDevServerPublicIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevServerPublicIpRequest struct{}"
	}

	return strings.Join([]string{"ListDevServerPublicIpRequest", string(data)}, " ")
}
