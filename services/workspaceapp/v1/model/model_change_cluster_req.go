package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeClusterReq 切换文件夹归属集群请求体
type ChangeClusterReq struct {

	// 目标集群ID
	TargetClusterGroupId string `json:"target_cluster_group_id"`

	// 目标项目配置ID
	TargetProjectConfigId string `json:"target_project_config_id"`

	// 文件系统id,数量区间 [1, 50]。
	CloudStorageAssignmentIds []string `json:"cloud_storage_assignment_ids"`
}

func (o ChangeClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeClusterReq struct{}"
	}

	return strings.Join([]string{"ChangeClusterReq", string(data)}, " ")
}
