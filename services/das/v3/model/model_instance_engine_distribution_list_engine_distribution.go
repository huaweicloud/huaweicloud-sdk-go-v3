package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceEngineDistributionListEngineDistribution struct {

	// 数据库类型
	DatastoreType *string `json:"datastore_type,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 实例信息
	InstanceInfos *[]InstanceEngineDistributionListInstanceInfos `json:"instance_infos,omitempty"`
}

func (o InstanceEngineDistributionListEngineDistribution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceEngineDistributionListEngineDistribution struct{}"
	}

	return strings.Join([]string{"InstanceEngineDistributionListEngineDistribution", string(data)}, " ")
}
