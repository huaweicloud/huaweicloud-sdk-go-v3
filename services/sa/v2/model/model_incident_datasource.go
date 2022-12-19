package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据源信息
type IncidentDatasource struct {

	// current page count
	SourceType *int32 `json:"source_type,omitempty"`

	// Id value
	DomainId *string `json:"domain_id,omitempty"`

	// Id value
	ProjectId *string `json:"project_id,omitempty"`

	// Id value
	RegionId *string `json:"region_id,omitempty"`
}

func (o IncidentDatasource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentDatasource struct{}"
	}

	return strings.Join([]string{"IncidentDatasource", string(data)}, " ")
}
