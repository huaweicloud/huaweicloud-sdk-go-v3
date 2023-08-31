package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMrsVersionListResponse Response Object
type ShowMrsVersionListResponse struct {

	// 可用集群版本列表
	ClusterVersions *[]string `json:"cluster_versions,omitempty"`
	HttpStatusCode  int       `json:"-"`
}

func (o ShowMrsVersionListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMrsVersionListResponse struct{}"
	}

	return strings.Join([]string{"ShowMrsVersionListResponse", string(data)}, " ")
}
