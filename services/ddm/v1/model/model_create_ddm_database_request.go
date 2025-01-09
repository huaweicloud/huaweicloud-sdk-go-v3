package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDdmDatabaseRequest Request Object
type CreateDdmDatabaseRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateDdmDatabaseRequestBody `json:"body,omitempty"`
}

func (o CreateDdmDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDdmDatabaseRequest struct{}"
	}

	return strings.Join([]string{"CreateDdmDatabaseRequest", string(data)}, " ")
}
