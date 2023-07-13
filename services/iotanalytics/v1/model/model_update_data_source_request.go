package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDataSourceRequest Request Object
type UpdateDataSourceRequest struct {

	// 数据源id
	DatasourceId string `json:"datasource_id"`

	Body *UpdateDatasourceReqDto `json:"body,omitempty"`
}

func (o UpdateDataSourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataSourceRequest struct{}"
	}

	return strings.Join([]string{"UpdateDataSourceRequest", string(data)}, " ")
}
