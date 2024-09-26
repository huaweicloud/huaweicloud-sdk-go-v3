package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpRiverConfig 瑞数配置项
type HttpRiverConfig struct {

	// 瑞数站点ID
	SiteId *string `json:"site_id,omitempty"`

	// 瑞数站点名称
	SiteName *string `json:"site_name,omitempty"`

	// 连接超时（毫秒）
	ConnectTimeout *int32 `json:"connect_timeout,omitempty"`

	// 读超时（毫秒）
	ReadTimeout *int32 `json:"read_timeout,omitempty"`

	// 写超时（毫秒）
	SendTimeout *int32 `json:"send_timeout,omitempty"`
}

func (o HttpRiverConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpRiverConfig struct{}"
	}

	return strings.Join([]string{"HttpRiverConfig", string(data)}, " ")
}
