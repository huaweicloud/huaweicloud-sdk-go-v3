package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGaussMysqlDnsResponse Response Object
type CreateGaussMysqlDnsResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGaussMysqlDnsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGaussMysqlDnsResponse struct{}"
	}

	return strings.Join([]string{"CreateGaussMysqlDnsResponse", string(data)}, " ")
}
