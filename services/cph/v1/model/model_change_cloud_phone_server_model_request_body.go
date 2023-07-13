package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeCloudPhoneServerModelRequestBody struct {

	// 云手机服务器的唯一标识。只有特定的服务器才能操作变更规格
	ServerId string `json:"server_id"`

	// 目标云手机服务器规格，不超过64个字节。仅允许相同代系服务器之间的规格切换。
	ServerModelName string `json:"server_model_name"`

	// 目标云手机规格。要求与变更前云手机规格路数相同，与目标云手机服务器规格匹配。
	PhoneModelName string `json:"phone_model_name"`

	ExtendParam *ChangeCloudPhoneServerModelRequestBodyExtendParam `json:"extend_param,omitempty"`
}

func (o ChangeCloudPhoneServerModelRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCloudPhoneServerModelRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeCloudPhoneServerModelRequestBody", string(data)}, " ")
}
