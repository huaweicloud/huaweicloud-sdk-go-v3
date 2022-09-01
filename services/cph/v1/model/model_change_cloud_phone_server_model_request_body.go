package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeCloudPhoneServerModelRequestBody struct {

	// 云手机服务器的唯一标识。只有特定的服务器才能操作变更规格
	ServerId string `json:"server_id" xml:"server_id"`

	// 要变更为的目标云手机服务器规格，不超过64个字节。 当前只支持填写physical.rx1.xlarge.special
	ServerModelName string `json:"server_model_name" xml:"server_model_name"`

	ExtendParam *ChangeCloudPhoneServerModelRequestBodyExtendParam `json:"extend_param,omitempty" xml:"extend_param"`
}

func (o ChangeCloudPhoneServerModelRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCloudPhoneServerModelRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeCloudPhoneServerModelRequestBody", string(data)}, " ")
}
