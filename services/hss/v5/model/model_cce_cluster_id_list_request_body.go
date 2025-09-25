package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CceClusterIdListRequestBody CCE集群配置请求
type CceClusterIdListRequestBody struct {

	// 集群id列表
	ClusterIdList []string `json:"cluster_id_list"`

	// 查询类型，包含如下:     - image : 镜像风险     - baseline : 基线风险     - vul : 漏洞风险     - event : 入侵风险
	DetectType *string `json:"detect_type,omitempty"`
}

func (o CceClusterIdListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CceClusterIdListRequestBody struct{}"
	}

	return strings.Join([]string{"CceClusterIdListRequestBody", string(data)}, " ")
}
