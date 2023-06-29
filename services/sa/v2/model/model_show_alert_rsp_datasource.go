package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlertRspDatasource 数据源信息
type ShowAlertRspDatasource struct {

	// current page count
	SourceType *int32 `json:"source_type,omitempty"`

	// Id value
	DomainId *string `json:"domain_id,omitempty"`

	// Id value
	ProjectId *string `json:"project_id,omitempty"`

	// Id value
	RegionId *string `json:"region_id,omitempty"`
}

func (o ShowAlertRspDatasource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertRspDatasource struct{}"
	}

	return strings.Join([]string{"ShowAlertRspDatasource", string(data)}, " ")
}
