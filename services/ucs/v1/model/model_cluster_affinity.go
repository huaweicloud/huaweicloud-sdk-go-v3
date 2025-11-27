package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterAffinity struct {

	// 指定要选择的集群名称列表
	ClusterNames *[]string `json:"clusterNames,omitempty"`

	// 指定要排除的集群名称列表
	Exclude *[]string `json:"exclude,omitempty"`
}

func (o ClusterAffinity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterAffinity struct{}"
	}

	return strings.Join([]string{"ClusterAffinity", string(data)}, " ")
}
