package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaSimpleServer
type NovaSimpleServer struct {

	// 云服务器名称。
	Name string `json:"name"`

	// 云服务器唯一标识。
	Id string `json:"id"`

	// 云服务器相关快捷链接信息。
	Links []NovaLink `json:"links"`
}

func (o NovaSimpleServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaSimpleServer struct{}"
	}

	return strings.Join([]string{"NovaSimpleServer", string(data)}, " ")
}
