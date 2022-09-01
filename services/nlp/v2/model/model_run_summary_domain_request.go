package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunSummaryDomainRequest struct {
	Body *SummaryDomainReq `json:"body,omitempty" xml:"body"`
}

func (o RunSummaryDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSummaryDomainRequest struct{}"
	}

	return strings.Join([]string{"RunSummaryDomainRequest", string(data)}, " ")
}
