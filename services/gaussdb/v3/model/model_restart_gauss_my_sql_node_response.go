package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartGaussMySqlNodeResponse Response Object
type RestartGaussMySqlNodeResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartGaussMySqlNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartGaussMySqlNodeResponse struct{}"
	}

	return strings.Join([]string{"RestartGaussMySqlNodeResponse", string(data)}, " ")
}
