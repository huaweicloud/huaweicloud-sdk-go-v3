package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodePoolConfigurationDetailsRespBody 获取指定节点池配置参数列表返回体
type ShowNodePoolConfigurationDetailsRespBody struct {

	// 配置参数，由key/value组成。
	Kubelet *[]PackageOptions `json:"kubelet,omitempty"`
}

func (o ShowNodePoolConfigurationDetailsRespBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodePoolConfigurationDetailsRespBody struct{}"
	}

	return strings.Join([]string{"ShowNodePoolConfigurationDetailsRespBody", string(data)}, " ")
}
