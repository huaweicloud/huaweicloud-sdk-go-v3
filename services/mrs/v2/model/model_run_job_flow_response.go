package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunJobFlowResponse Response Object
type RunJobFlowResponse struct {

	// 集群创建成功后系统返回的集群ID值。
	ClusterId      *string `json:"cluster_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunJobFlowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunJobFlowResponse struct{}"
	}

	return strings.Join([]string{"RunJobFlowResponse", string(data)}, " ")
}
