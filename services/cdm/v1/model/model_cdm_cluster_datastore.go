package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CdmClusterDatastore struct {

	// 服务ID，用于区分不同服务。
	Id string `json:"id"`

	// 服务名称。
	Name string `json:"name"`

	// 是否支持大规格集群。
	BigclusterEnable bool `json:"bigclusterEnable"`

	// 默认版本。
	DefaultVersion string `json:"defaultVersion"`

	// 版本。
	Versions []CdmClusterVersion `json:"versions"`
}

func (o CdmClusterDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmClusterDatastore struct{}"
	}

	return strings.Join([]string{"CdmClusterDatastore", string(data)}, " ")
}
