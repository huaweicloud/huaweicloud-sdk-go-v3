package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GrantAccessServiceRequestBody 接入服务授权信息列表
type GrantAccessServiceRequestBody struct {

	// 接入服务授权信息列表
	AccessServices []AccessServiceInfo `json:"access_services"`
}

func (o GrantAccessServiceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GrantAccessServiceRequestBody struct{}"
	}

	return strings.Join([]string{"GrantAccessServiceRequestBody", string(data)}, " ")
}
