package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListInstanceMultiNodesSingleMetricInstanceInfos struct {

	// 数据库类型
	DatastoreType string `json:"datastore_type"`

	// 节点列表
	NodeInfos []ListInstanceMultiNodesSingleMetricNodeInfos `json:"node_infos"`
}

func (o ListInstanceMultiNodesSingleMetricInstanceInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceMultiNodesSingleMetricInstanceInfos struct{}"
	}

	return strings.Join([]string{"ListInstanceMultiNodesSingleMetricInstanceInfos", string(data)}, " ")
}
