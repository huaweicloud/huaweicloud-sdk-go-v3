package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompleteResourceTokenAuthRequest Request Object
type CompleteResourceTokenAuthRequest struct {
	Body *CompleteResourceTokenAuthRequestBody `json:"body,omitempty"`
}

func (o CompleteResourceTokenAuthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompleteResourceTokenAuthRequest struct{}"
	}

	return strings.Join([]string{"CompleteResourceTokenAuthRequest", string(data)}, " ")
}
