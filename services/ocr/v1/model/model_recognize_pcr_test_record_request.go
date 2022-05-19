package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizePcrTestRecordRequest struct {
	Body *PcrTestRecordRequestBody `json:"body,omitempty"`
}

func (o RecognizePcrTestRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizePcrTestRecordRequest struct{}"
	}

	return strings.Join([]string{"RecognizePcrTestRecordRequest", string(data)}, " ")
}
