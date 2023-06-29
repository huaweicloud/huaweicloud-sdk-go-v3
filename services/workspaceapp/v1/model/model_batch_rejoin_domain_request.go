package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRejoinDomainRequest Request Object
type BatchRejoinDomainRequest struct {
	Body *BatchRejoinDomainReq `json:"body,omitempty"`
}

func (o BatchRejoinDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRejoinDomainRequest struct{}"
	}

	return strings.Join([]string{"BatchRejoinDomainRequest", string(data)}, " ")
}
