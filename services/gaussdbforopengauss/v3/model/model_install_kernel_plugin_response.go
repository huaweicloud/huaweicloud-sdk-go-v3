package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstallKernelPluginResponse Response Object
type InstallKernelPluginResponse struct {

	// 插件安装工作流id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o InstallKernelPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallKernelPluginResponse struct{}"
	}

	return strings.Join([]string{"InstallKernelPluginResponse", string(data)}, " ")
}
