package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDisassociatePublicipsRequest Request Object
type BatchDisassociatePublicipsRequest struct {
	Body *BatchDeletePublicIpRequestBody `json:"body,omitempty"`
}

func (o BatchDisassociatePublicipsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisassociatePublicipsRequest struct{}"
	}

	return strings.Join([]string{"BatchDisassociatePublicipsRequest", string(data)}, " ")
}
