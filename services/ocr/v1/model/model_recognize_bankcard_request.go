package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeBankcardRequest struct {
	Body *BankcardRequestBody `json:"body,omitempty"`
}

func (o RecognizeBankcardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeBankcardRequest struct{}"
	}

	return strings.Join([]string{"RecognizeBankcardRequest", string(data)}, " ")
}
