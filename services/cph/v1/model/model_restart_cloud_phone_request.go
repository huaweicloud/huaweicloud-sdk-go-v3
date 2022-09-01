package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RestartCloudPhoneRequest struct {
	Body *ResetRestartRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RestartCloudPhoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartCloudPhoneRequest struct{}"
	}

	return strings.Join([]string{"RestartCloudPhoneRequest", string(data)}, " ")
}
