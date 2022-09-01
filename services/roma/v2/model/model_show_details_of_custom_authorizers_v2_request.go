package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDetailsOfCustomAuthorizersV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 自定义认证的编号
	AuthorizerId string `json:"authorizer_id" xml:"authorizer_id"`
}

func (o ShowDetailsOfCustomAuthorizersV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfCustomAuthorizersV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfCustomAuthorizersV2Request", string(data)}, " ")
}
