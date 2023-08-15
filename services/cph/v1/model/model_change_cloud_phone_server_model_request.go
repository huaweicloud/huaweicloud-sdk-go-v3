package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeCloudPhoneServerModelRequest Request Object
type ChangeCloudPhoneServerModelRequest struct {
	Body *ChangeCloudPhoneServerModelRequestBody `json:"body,omitempty"`
}

func (o ChangeCloudPhoneServerModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCloudPhoneServerModelRequest struct{}"
	}

	return strings.Join([]string{"ChangeCloudPhoneServerModelRequest", string(data)}, " ")
}
