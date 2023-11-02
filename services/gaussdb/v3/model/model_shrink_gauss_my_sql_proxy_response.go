package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShrinkGaussMySqlProxyResponse Response Object
type ShrinkGaussMySqlProxyResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShrinkGaussMySqlProxyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkGaussMySqlProxyResponse struct{}"
	}

	return strings.Join([]string{"ShrinkGaussMySqlProxyResponse", string(data)}, " ")
}
