package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDdmDatabaseResponse Response Object
type CreateDdmDatabaseResponse struct {

	// 逻辑库名
	Name *string `json:"name,omitempty"`

	// 工作流id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDdmDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDdmDatabaseResponse struct{}"
	}

	return strings.Join([]string{"CreateDdmDatabaseResponse", string(data)}, " ")
}
