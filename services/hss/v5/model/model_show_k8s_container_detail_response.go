package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowK8sContainerDetailResponse Response Object
type ShowK8sContainerDetailResponse struct {

	// 服务名称
	ServiceName *string `json:"service_name,omitempty"`

	// 服务用户名
	ServiceUsername *string `json:"service_username,omitempty"`

	// 服务密码
	ServicePassword *string `json:"service_password,omitempty"`

	// 容器各服务端口信息
	ServicePortList *[]ServicePortInfo `json:"service_port_list,omitempty"`

	// 数据仿真，默认关闭。开启后将在沙箱中注入仿真数据 - open : 开启 - close : 关闭
	EnableSimulate *string `json:"enable_simulate,omitempty"`

	// 沙箱域名，域名之间使用 ',' 隔开
	Hosts *string `json:"hosts,omitempty"`

	Extra          *ContainerExtraInfo `json:"extra,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowK8sContainerDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowK8sContainerDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowK8sContainerDetailResponse", string(data)}, " ")
}
