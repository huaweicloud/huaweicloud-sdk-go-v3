package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVoiceRequest Request Object
type CreateVoiceRequest struct {
	Body *RegisterVoiceReq `json:"body,omitempty"`
}

func (o CreateVoiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVoiceRequest struct{}"
	}

	return strings.Join([]string{"CreateVoiceRequest", string(data)}, " ")
}
