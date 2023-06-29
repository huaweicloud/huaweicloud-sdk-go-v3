package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunDomainSentimentRequest Request Object
type RunDomainSentimentRequest struct {
	Body *DomainSentimentReq `json:"body,omitempty"`
}

func (o RunDomainSentimentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDomainSentimentRequest struct{}"
	}

	return strings.Join([]string{"RunDomainSentimentRequest", string(data)}, " ")
}
