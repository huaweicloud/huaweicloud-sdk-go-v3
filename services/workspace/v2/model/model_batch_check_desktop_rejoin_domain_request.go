package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCheckDesktopRejoinDomainRequest Request Object
type BatchCheckDesktopRejoinDomainRequest struct {
	Body *BatchCheckRejoinDomainReq `json:"body,omitempty"`
}

func (o BatchCheckDesktopRejoinDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckDesktopRejoinDomainRequest struct{}"
	}

	return strings.Join([]string{"BatchCheckDesktopRejoinDomainRequest", string(data)}, " ")
}
