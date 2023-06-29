package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDatasourceInfoRequest Request Object
type UpdateDatasourceInfoRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据源ID
	DatasourceId string `json:"datasource_id"`

	Body *DatasourceInfo `json:"body,omitempty"`
}

func (o UpdateDatasourceInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatasourceInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateDatasourceInfoRequest", string(data)}, " ")
}
