package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetHyperinstanceGroupRequest Request Object
type GetHyperinstanceGroupRequest struct {

	// **参数解释**：需查询的hyperinstance-group ID。 **默认取值**：不涉及。
	Id string `json:"id"`
}

func (o GetHyperinstanceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetHyperinstanceGroupRequest struct{}"
	}

	return strings.Join([]string{"GetHyperinstanceGroupRequest", string(data)}, " ")
}
