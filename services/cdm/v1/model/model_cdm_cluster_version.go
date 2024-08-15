package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CdmClusterVersion 版本信息。
type CdmClusterVersion struct {

	// 版本状态。
	Active *string `json:"active,omitempty"`

	// 版本ID。
	Id *string `json:"id,omitempty"`

	// 版本镜像。
	Image *string `json:"image,omitempty"`

	// 版本名称。
	Name *string `json:"name,omitempty"`

	// 版本的包。
	Packages *string `json:"packages,omitempty"`

	// 服务ID，用于区分不同服务。
	Datastore *string `json:"datastore,omitempty"`

	// 链接信息。
	Links *[]ClusterLinks `json:"links,omitempty"`
}

func (o CdmClusterVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmClusterVersion struct{}"
	}

	return strings.Join([]string{"CdmClusterVersion", string(data)}, " ")
}
