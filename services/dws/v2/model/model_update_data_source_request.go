package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDataSourceRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 数据源id
	ExtDataSourceId string `json:"ext_data_source_id"`

	Body *ReconfigureExtDataSourceAction `json:"body,omitempty"`
}

func (o UpdateDataSourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataSourceRequest struct{}"
	}

	return strings.Join([]string{"UpdateDataSourceRequest", string(data)}, " ")
}
