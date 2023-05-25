package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMetadataResponse struct {

	// 备份ID
	BackupId *string `json:"backup_id,omitempty"`

	// 云服务器备份信息
	Backups *string `json:"backups,omitempty"`

	// 云服务器规格信息
	Flavor *string `json:"flavor,omitempty"`

	// 云服务器浮动IP信息
	Floatingips *[]string `json:"floatingips,omitempty"`

	// 云服务器接口信息
	Interface *string `json:"interface,omitempty"`

	// 云服务器端口信息
	Ports *[]string `json:"ports,omitempty"`

	// 云服务器信息
	Server *string `json:"server,omitempty"`

	// 云服务器卷信息
	Volumes        *[]string `json:"volumes,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetadataResponse struct{}"
	}

	return strings.Join([]string{"ShowMetadataResponse", string(data)}, " ")
}
