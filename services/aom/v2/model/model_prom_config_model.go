package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PromConfigModel struct {

	// 普罗实例remote-write地址
	RemoteWriteUrl *string `json:"remote_write_url,omitempty"`

	// 普罗实例remote-read地址
	RemoteReadUrl *string `json:"remote_read_url,omitempty"`

	// 普罗实例调用url
	PromHttpApiEndpoint *string `json:"prom_http_api_endpoint,omitempty"`

	// 普罗实例关联dashboard的dashboard ID
	DashboardId *string `json:"dashboard_id,omitempty"`

	// 普罗实例所属的region
	RegionId *string `json:"region_id,omitempty"`
}

func (o PromConfigModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PromConfigModel struct{}"
	}

	return strings.Join([]string{"PromConfigModel", string(data)}, " ")
}
