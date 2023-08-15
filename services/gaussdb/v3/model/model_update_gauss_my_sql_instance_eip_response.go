package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGaussMySqlInstanceEipResponse Response Object
type UpdateGaussMySqlInstanceEipResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGaussMySqlInstanceEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlInstanceEipResponse struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlInstanceEipResponse", string(data)}, " ")
}
