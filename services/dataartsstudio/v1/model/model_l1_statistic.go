package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type L1Statistic struct {

	// 主题域分组名称
	SubjectAreaGroupName *string `json:"subject_area_group_name,omitempty"`

	// 主题域分组英文名称
	SubjectAreaGroupNameEn *string `json:"subject_area_group_name_en,omitempty"`

	// 主题域分组的guid
	SubjectAreaGroupGuid *string `json:"subject_area_group_guid,omitempty"`

	// 主题域分组序号
	Ordinal *int32 `json:"ordinal,omitempty"`

	// 主题总数
	SubjectAreaCount *int32 `json:"subject_area_count,omitempty"`

	// 业务对象总数
	BusinessObjectCount *int32 `json:"business_object_count,omitempty"`

	// 逻辑实体总数
	LogicEntityCount *int32 `json:"logic_entity_count,omitempty"`

	// 主题统计信息
	SubjectAreaStatistics *[]L2Statistic `json:"subject_area_statistics,omitempty"`
}

func (o L1Statistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L1Statistic struct{}"
	}

	return strings.Join([]string{"L1Statistic", string(data)}, " ")
}
