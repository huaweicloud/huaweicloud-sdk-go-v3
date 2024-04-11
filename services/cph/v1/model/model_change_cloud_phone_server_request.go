package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeCloudPhoneServerRequest Request Object
type ChangeCloudPhoneServerRequest struct {

	// 云手机服务器的唯一标识。
	ServerId string `json:"server_id"`

	Body *ChangeCloudPhoneServerRequestBody `json:"body,omitempty"`
}

func (o ChangeCloudPhoneServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCloudPhoneServerRequest struct{}"
	}

	return strings.Join([]string{"ChangeCloudPhoneServerRequest", string(data)}, " ")
}
