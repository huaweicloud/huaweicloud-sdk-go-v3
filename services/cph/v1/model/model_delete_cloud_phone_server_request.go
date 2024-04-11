package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCloudPhoneServerRequest Request Object
type DeleteCloudPhoneServerRequest struct {
	Body *DeleteCloudPhoneServerRequestBody `json:"body,omitempty"`
}

func (o DeleteCloudPhoneServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCloudPhoneServerRequest struct{}"
	}

	return strings.Join([]string{"DeleteCloudPhoneServerRequest", string(data)}, " ")
}
