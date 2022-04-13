package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDatasourceInfoRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *DatasourceInfo `json:"body,omitempty"`
}

func (o CreateDatasourceInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatasourceInfoRequest struct{}"
	}

	return strings.Join([]string{"CreateDatasourceInfoRequest", string(data)}, " ")
}
