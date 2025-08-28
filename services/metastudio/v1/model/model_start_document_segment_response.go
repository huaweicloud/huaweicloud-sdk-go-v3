package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartDocumentSegmentResponse Response Object
type StartDocumentSegmentResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartDocumentSegmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartDocumentSegmentResponse struct{}"
	}

	return strings.Join([]string{"StartDocumentSegmentResponse", string(data)}, " ")
}
