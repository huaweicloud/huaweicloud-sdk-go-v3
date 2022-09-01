package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ResetCloudPhoneRequest struct {
	Body *ResetRestartRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ResetCloudPhoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetCloudPhoneRequest struct{}"
	}

	return strings.Join([]string{"ResetCloudPhoneRequest", string(data)}, " ")
}
