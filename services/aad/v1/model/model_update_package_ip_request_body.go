package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePackageIpRequestBody 更新全量防护对象的请求体
type UpdatePackageIpRequestBody struct {

	// 全量防护ip列表
	ProtectedIpList []UpdateProtectedIpInPolicyBody `json:"protected_ip_list"`
}

func (o UpdatePackageIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePackageIpRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePackageIpRequestBody", string(data)}, " ")
}
