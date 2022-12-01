package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDataSourceRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	Body *interface{} `json:"body,omitempty"`
}

func (o CreateDataSourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataSourceRequest struct{}"
	}

	return strings.Join([]string{"CreateDataSourceRequest", string(data)}, " ")
}
