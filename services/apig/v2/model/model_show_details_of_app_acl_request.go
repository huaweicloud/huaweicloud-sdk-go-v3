package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailsOfAppAclRequest Request Object
type ShowDetailsOfAppAclRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 凭据编号
	AppId string `json:"app_id"`
}

func (o ShowDetailsOfAppAclRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfAppAclRequest struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfAppAclRequest", string(data)}, " ")
}
