package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDdmDatabaseResponse Response Object
type DeleteDdmDatabaseResponse struct {

	// 工作流id。
	JobId *string `json:"job_id,omitempty"`

	// 逻辑库名
	DatabaseName   *string `json:"database_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDdmDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDdmDatabaseResponse struct{}"
	}

	return strings.Join([]string{"DeleteDdmDatabaseResponse", string(data)}, " ")
}
