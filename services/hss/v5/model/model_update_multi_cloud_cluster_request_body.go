package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMultiCloudClusterRequestBody 更新的集群信息
type UpdateMultiCloudClusterRequestBody struct {

	// 当前有效期结束时间
	CurrentExpirationDate int64 `json:"current_expiration_date"`

	// 是否升级代理版本
	Upgrade *bool `json:"upgrade,omitempty"`

	// kubeconfig文件
	Kubeconfig *string `json:"kubeconfig,omitempty"`
}

func (o UpdateMultiCloudClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMultiCloudClusterRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateMultiCloudClusterRequestBody", string(data)}, " ")
}
