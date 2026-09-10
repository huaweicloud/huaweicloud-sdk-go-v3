package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterResponse Response Object
type CreateClusterResponse struct {

	// 请求执行是否成功。\"true\"表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 新增集群的名称。
	ClusterName *string `json:"cluster_name,omitempty"`

	// 异步作业id，用于查询作业状态。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterResponse", string(data)}, " ")
}
