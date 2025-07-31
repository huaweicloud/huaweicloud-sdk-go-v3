package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMultiCloudClusterProxyScriptResponse Response Object
type ShowMultiCloudClusterProxyScriptResponse struct {

	// 代理安装脚本
	InstallScript  *string `json:"install_script,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMultiCloudClusterProxyScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMultiCloudClusterProxyScriptResponse struct{}"
	}

	return strings.Join([]string{"ShowMultiCloudClusterProxyScriptResponse", string(data)}, " ")
}
