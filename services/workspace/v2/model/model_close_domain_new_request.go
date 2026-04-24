package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloseDomainNewRequest Request Object
type CloseDomainNewRequest struct {
	Body *CloseDomainNewReq `json:"body,omitempty"`
}

func (o CloseDomainNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloseDomainNewRequest struct{}"
	}

	return strings.Join([]string{"CloseDomainNewRequest", string(data)}, " ")
}
