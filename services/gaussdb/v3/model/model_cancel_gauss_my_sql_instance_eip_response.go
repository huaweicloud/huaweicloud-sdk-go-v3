package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelGaussMySqlInstanceEipResponse Response Object
type CancelGaussMySqlInstanceEipResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelGaussMySqlInstanceEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelGaussMySqlInstanceEipResponse struct{}"
	}

	return strings.Join([]string{"CancelGaussMySqlInstanceEipResponse", string(data)}, " ")
}
