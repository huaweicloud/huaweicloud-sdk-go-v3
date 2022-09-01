package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ImportDataResponse struct {

	// 表ID。
	TableId *string `json:"table_id,omitempty" xml:"table_id"`

	// 作业ID。
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 作业运行ID。
	RunId          *string `json:"run_id,omitempty" xml:"run_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportDataResponse struct{}"
	}

	return strings.Join([]string{"ImportDataResponse", string(data)}, " ")
}
