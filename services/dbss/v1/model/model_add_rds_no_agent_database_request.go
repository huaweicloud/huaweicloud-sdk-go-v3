package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddRdsNoAgentDatabaseRequest Request Object
type AddRdsNoAgentDatabaseRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *RdsNoAgentDbRequest `json:"body,omitempty"`
}

func (o AddRdsNoAgentDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRdsNoAgentDatabaseRequest struct{}"
	}

	return strings.Join([]string{"AddRdsNoAgentDatabaseRequest", string(data)}, " ")
}
