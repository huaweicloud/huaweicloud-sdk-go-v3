package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicationEndpoint struct {

	// 创建application的时间 时间格式为UTC时间，YYYY-MM-DDTHH:MM:SSZ。
	CreateTime string `json:"create_time" xml:"create_time"`

	// Application endpoint的唯一资源标识。
	EndpointUrn string `json:"endpoint_urn" xml:"endpoint_urn"`

	// 用户自定义数据 最大长度支持UTF-8编码后2048字节。
	UserData string `json:"user_data" xml:"user_data"`

	// endpoint启用开关 true或false字符串。
	Enabled string `json:"enabled" xml:"enabled"`

	// 设备token 最大长度512个字节。
	Token string `json:"token" xml:"token"`
}

func (o ApplicationEndpoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationEndpoint struct{}"
	}

	return strings.Join([]string{"ApplicationEndpoint", string(data)}, " ")
}
