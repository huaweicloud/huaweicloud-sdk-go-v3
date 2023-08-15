package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGaussMySqlInstancePortResponse Response Object
type UpdateGaussMySqlInstancePortResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGaussMySqlInstancePortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlInstancePortResponse struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlInstancePortResponse", string(data)}, " ")
}
