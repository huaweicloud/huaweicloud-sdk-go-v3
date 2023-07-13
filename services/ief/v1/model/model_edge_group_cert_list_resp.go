package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EdgeGroupCertListResp 边缘节点组证书列表返回体
type EdgeGroupCertListResp struct {

	// 边缘节点组证书数目
	Count *int32 `json:"count,omitempty"`

	// 边缘节点证书详情
	Groupcerts *[]EdgeGroupCert `json:"groupcerts,omitempty"`
}

func (o EdgeGroupCertListResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeGroupCertListResp struct{}"
	}

	return strings.Join([]string{"EdgeGroupCertListResp", string(data)}, " ")
}
