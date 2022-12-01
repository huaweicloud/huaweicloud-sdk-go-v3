package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建集群校验请求体
type ClusterCheckRequestBody struct {
	Cluster *ClusterCheckBody `json:"cluster"`
}

func (o ClusterCheckRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterCheckRequestBody struct{}"
	}

	return strings.Join([]string{"ClusterCheckRequestBody", string(data)}, " ")
}
