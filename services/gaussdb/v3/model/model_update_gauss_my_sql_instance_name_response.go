package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateGaussMySqlInstanceNameResponse struct {
	// 修改实例名称的任务id

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGaussMySqlInstanceNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlInstanceNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlInstanceNameResponse", string(data)}, " ")
}
