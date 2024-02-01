package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartRdSforMysqlProxyResponse Response Object
type RestartRdSforMysqlProxyResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartRdSforMysqlProxyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartRdSforMysqlProxyResponse struct{}"
	}

	return strings.Join([]string{"RestartRdSforMysqlProxyResponse", string(data)}, " ")
}
