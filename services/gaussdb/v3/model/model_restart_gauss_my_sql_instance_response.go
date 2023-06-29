package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartGaussMySqlInstanceResponse Response Object
type RestartGaussMySqlInstanceResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartGaussMySqlInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartGaussMySqlInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestartGaussMySqlInstanceResponse", string(data)}, " ")
}
