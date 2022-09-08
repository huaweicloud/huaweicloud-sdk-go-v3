package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePhoneNameRequest struct {

	// 云手机id。
	PhoneId string `json:"phone_id"`

	Body *UpdatePhoneNameRequestBody `json:"body,omitempty"`
}

func (o UpdatePhoneNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePhoneNameRequest struct{}"
	}

	return strings.Join([]string{"UpdatePhoneNameRequest", string(data)}, " ")
}
