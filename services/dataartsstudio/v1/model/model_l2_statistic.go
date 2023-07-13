package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type L2Statistic struct {

	// 主题名称
	SubjectAreaName *string `json:"subject_area_name,omitempty"`

	// 主题的guid
	SubjectAreaGuid *string `json:"subject_area_guid,omitempty"`

	// 主题序号
	Ordinal *int32 `json:"ordinal,omitempty"`

	// 业务对象总数
	BusinessObjectCount *int32 `json:"business_object_count,omitempty"`

	// 逻辑实体总数
	LogicEntityCount *int32 `json:"logic_entity_count,omitempty"`
}

func (o L2Statistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L2Statistic struct{}"
	}

	return strings.Join([]string{"L2Statistic", string(data)}, " ")
}
