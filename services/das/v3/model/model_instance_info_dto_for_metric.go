package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceInfoDtoForMetric 实例信息
type InstanceInfoDtoForMetric struct {

	// 数据库类型
	DatastoreType string `json:"datastore_type"`

	// 实例信息列表
	NodeInfos []NodeInfoForMetric `json:"node_infos"`
}

func (o InstanceInfoDtoForMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceInfoDtoForMetric struct{}"
	}

	return strings.Join([]string{"InstanceInfoDtoForMetric", string(data)}, " ")
}
