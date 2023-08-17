package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyGaussMysqlDnsResponse Response Object
type ModifyGaussMysqlDnsResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyGaussMysqlDnsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyGaussMysqlDnsResponse struct{}"
	}

	return strings.Join([]string{"ModifyGaussMysqlDnsResponse", string(data)}, " ")
}
