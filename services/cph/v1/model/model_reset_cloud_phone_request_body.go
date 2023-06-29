package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetCloudPhoneRequestBody 重置云手机请求体。
type ResetCloudPhoneRequestBody struct {

	// 云手机镜像
	ImageId *string `json:"image_id,omitempty"`

	// 云手机列表
	Phones []PhoneProperty `json:"phones"`
}

func (o ResetCloudPhoneRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetCloudPhoneRequestBody struct{}"
	}

	return strings.Join([]string{"ResetCloudPhoneRequestBody", string(data)}, " ")
}
