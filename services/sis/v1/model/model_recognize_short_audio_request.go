package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeShortAudioRequest struct {
	Body *PostShortAudioReq `json:"body,omitempty" xml:"body"`
}

func (o RecognizeShortAudioRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeShortAudioRequest struct{}"
	}

	return strings.Join([]string{"RecognizeShortAudioRequest", string(data)}, " ")
}
