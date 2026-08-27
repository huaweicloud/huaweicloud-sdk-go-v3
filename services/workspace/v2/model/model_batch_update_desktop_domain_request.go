package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateDesktopDomainRequest Request Object
type BatchUpdateDesktopDomainRequest struct {
	Body *BatchRejoinDomainReq `json:"body,omitempty"`
}

func (o BatchUpdateDesktopDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateDesktopDomainRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateDesktopDomainRequest", string(data)}, " ")
}
