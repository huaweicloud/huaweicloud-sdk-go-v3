package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DirectionalConnection 有向连接
type DirectionalConnection struct {

	// 实例名称。
	Name *string `json:"name,omitempty"`

	// 实例ID。
	Id string `json:"id"`

	// 功能说明：本端接入点的编码。
	LocalSiteCode *string `json:"local_site_code,omitempty"`

	// 功能说明：远端接入点的编码。
	RemoteSiteCode *string `json:"remote_site_code,omitempty"`
}

func (o DirectionalConnection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DirectionalConnection struct{}"
	}

	return strings.Join([]string{"DirectionalConnection", string(data)}, " ")
}
