package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetCloudPhoneRequest Request Object
type ResetCloudPhoneRequest struct {
	Body *ResetCloudPhoneRequestBody `json:"body,omitempty"`
}

func (o ResetCloudPhoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetCloudPhoneRequest struct{}"
	}

	return strings.Join([]string{"ResetCloudPhoneRequest", string(data)}, " ")
}
