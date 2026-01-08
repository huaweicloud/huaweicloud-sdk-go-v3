package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteLoadbalancersRequest Request Object
type BatchDeleteLoadbalancersRequest struct {
	Body *BatchDeleteLoadbalancersRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteLoadbalancersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteLoadbalancersRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteLoadbalancersRequest", string(data)}, " ")
}
