package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstanceCrossVpcIpReq struct {
	// 用户自定义的advertised_ip_contents键值对。  键是listeners IP。  值是advertised.listeners IP，或者域名。  > IP修改未修改项也需填上。

	AdvertisedIpContents map[string]string `json:"advertised_ip_contents"`
}

func (o UpdateInstanceCrossVpcIpReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceCrossVpcIpReq struct{}"
	}

	return strings.Join([]string{"UpdateInstanceCrossVpcIpReq", string(data)}, " ")
}
