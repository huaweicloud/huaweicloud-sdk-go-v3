package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartCloudPhoneServerRequest Request Object
type RestartCloudPhoneServerRequest struct {
	Body *RestartCloudPhoneServerRequestBody `json:"body,omitempty"`
}

func (o RestartCloudPhoneServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartCloudPhoneServerRequest struct{}"
	}

	return strings.Join([]string{"RestartCloudPhoneServerRequest", string(data)}, " ")
}
