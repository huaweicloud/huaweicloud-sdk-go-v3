package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAddressSetsRequest Request Object
type BatchDeleteAddressSetsRequest struct {
	Body *BatchDeleteAddressSetsDto `json:"body,omitempty"`
}

func (o BatchDeleteAddressSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAddressSetsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteAddressSetsRequest", string(data)}, " ")
}
