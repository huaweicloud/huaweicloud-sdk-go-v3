package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterRequestBody 创建集群发起的请求的请求体对象。
type CreateClusterRequestBody struct {
	Cluster *Cluster `json:"cluster"`
}

func (o CreateClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterRequestBody struct{}"
	}

	return strings.Join([]string{"CreateClusterRequestBody", string(data)}, " ")
}
