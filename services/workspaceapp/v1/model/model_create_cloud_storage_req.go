package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudStorageReq 创建项目配置关联。
type CreateCloudStorageReq struct {

	// project_config_id,数量区间 [1, 50]。
	ProjectConfigIds []string `json:"project_config_ids"`
}

func (o CreateCloudStorageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudStorageReq struct{}"
	}

	return strings.Join([]string{"CreateCloudStorageReq", string(data)}, " ")
}
