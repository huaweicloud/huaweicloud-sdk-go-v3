package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 重启云手机请求体。
type RestartCloudPhoneRequestBody struct {

	// 云手机镜像
	ImageId *string `json:"image_id,omitempty"`

	// 云手机列表
	Phones []PhoneProperty `json:"phones"`
}

func (o RestartCloudPhoneRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartCloudPhoneRequestBody struct{}"
	}

	return strings.Join([]string{"RestartCloudPhoneRequestBody", string(data)}, " ")
}
