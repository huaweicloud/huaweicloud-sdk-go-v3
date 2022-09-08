package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizePcrTestRecordResponse struct {
	Result         *PcrTestRecordResult `json:"result,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o RecognizePcrTestRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizePcrTestRecordResponse struct{}"
	}

	return strings.Join([]string{"RecognizePcrTestRecordResponse", string(data)}, " ")
}
