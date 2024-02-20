package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartResourceScanRequest Request Object
type StartResourceScanRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	Body *StartResourceScanReqBody `json:"body,omitempty"`
}

func (o StartResourceScanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartResourceScanRequest struct{}"
	}

	return strings.Join([]string{"StartResourceScanRequest", string(data)}, " ")
}
