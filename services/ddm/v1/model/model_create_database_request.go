package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDatabaseRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *CreateDatabaseReq `json:"body,omitempty" xml:"body"`
}

func (o CreateDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseRequest struct{}"
	}

	return strings.Join([]string{"CreateDatabaseRequest", string(data)}, " ")
}
