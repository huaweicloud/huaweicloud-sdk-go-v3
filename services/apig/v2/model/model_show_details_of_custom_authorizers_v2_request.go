package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailsOfCustomAuthorizersV2Request Request Object
type ShowDetailsOfCustomAuthorizersV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 自定义认证的编号
	AuthorizerId string `json:"authorizer_id"`
}

func (o ShowDetailsOfCustomAuthorizersV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfCustomAuthorizersV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfCustomAuthorizersV2Request", string(data)}, " ")
}
