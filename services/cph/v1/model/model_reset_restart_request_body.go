package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResetRestartRequestBody struct {

	// 云手机镜像
	ImageId *string `json:"image_id,omitempty" xml:"image_id"`

	// 云手机列表
	Phones []ResetRestartRequestBodyPhones `json:"phones" xml:"phones"`
}

func (o ResetRestartRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetRestartRequestBody struct{}"
	}

	return strings.Join([]string{"ResetRestartRequestBody", string(data)}, " ")
}
