package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectivityTestRequest Request Object
type CreateConnectivityTestRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateConnectivityTestRequestBody `json:"body,omitempty"`
}

func (o CreateConnectivityTestRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectivityTestRequest struct{}"
	}

	return strings.Join([]string{"CreateConnectivityTestRequest", string(data)}, " ")
}
