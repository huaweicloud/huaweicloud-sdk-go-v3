package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHyperinstanceGroupRequest Request Object
type DeleteHyperinstanceGroupRequest struct {

	// **参数解释**：要删除的hyperinstance-group ID。 **默认取值**：不涉及。
	Id string `json:"id"`
}

func (o DeleteHyperinstanceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHyperinstanceGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteHyperinstanceGroupRequest", string(data)}, " ")
}
