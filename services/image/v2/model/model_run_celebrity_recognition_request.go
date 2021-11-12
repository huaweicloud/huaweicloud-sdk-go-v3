package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunCelebrityRecognitionRequest struct {
	Body *CelebrityRecognitionReq `json:"body,omitempty"`
}

func (o RunCelebrityRecognitionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCelebrityRecognitionRequest struct{}"
	}

	return strings.Join([]string{"RunCelebrityRecognitionRequest", string(data)}, " ")
}
