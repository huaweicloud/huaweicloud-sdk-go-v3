package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudStorageReq 创建项目配置关联。
type CreateCloudStorageReq struct {

	// 创建项目配置关联ID列表。
	ProjectConfigClusterGroupIdList []ProjectConfigClusterGroupIdEntity `json:"project_config_cluster_group_id_list"`
}

func (o CreateCloudStorageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudStorageReq struct{}"
	}

	return strings.Join([]string{"CreateCloudStorageReq", string(data)}, " ")
}
