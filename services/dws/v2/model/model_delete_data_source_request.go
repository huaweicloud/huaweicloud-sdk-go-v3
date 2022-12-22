package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDataSourceRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 数据源id
	ExtDataSourceId string `json:"ext_data_source_id"`
}

func (o DeleteDataSourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataSourceRequest struct{}"
	}

	return strings.Join([]string{"DeleteDataSourceRequest", string(data)}, " ")
}
