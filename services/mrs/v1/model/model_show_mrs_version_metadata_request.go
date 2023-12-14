package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMrsVersionMetadataRequest Request Object
type ShowMrsVersionMetadataRequest struct {

	// 集群版本。例如“MRS 3.1.0”。如果请求客户端不支持自动转义，则需要将空格转义为%20，例如“MRS%203.1.0”。
	VersionName string `json:"version_name"`

	// 集群ID。如果指定集群ID，则获取该集群做过补丁更新的最新版本元数据。
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o ShowMrsVersionMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMrsVersionMetadataRequest struct{}"
	}

	return strings.Join([]string{"ShowMrsVersionMetadataRequest", string(data)}, " ")
}
