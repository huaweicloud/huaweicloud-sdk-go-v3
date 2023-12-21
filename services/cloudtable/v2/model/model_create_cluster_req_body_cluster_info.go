package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterReqBodyClusterInfo 集群详情
type CreateClusterReqBodyClusterInfo struct {

	// 特性属性开关      * 属性开关必须以enable开头，value必须为true|false      * doris: enable_broker      * hbase: storage_io_type（COMMON，ULTRAHIGH，两种取值），enable_open_tsdb（默认false，若为true需要在集群节点信息列表中指定tsd节点个数），enable_broker      示例：      \"feature_map\":{\"enable_broker\":\"false\"}       \"feature_map\":{\"enable_lemon\":\"false\",\"enable_open_tsdb\":\"false\",\"storage_io_type\": \"COMMON\"}
	FeatureMap map[string]string `json:"feature_map"`

	// 集群节点信息类
	ClusterInstanceInfo []CreateClusterInstanceBody `json:"cluster_instance_info"`
}

func (o CreateClusterReqBodyClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterReqBodyClusterInfo struct{}"
	}

	return strings.Join([]string{"CreateClusterReqBodyClusterInfo", string(data)}, " ")
}
