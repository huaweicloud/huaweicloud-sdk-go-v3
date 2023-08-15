package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetGaussMySqlProxyWeightResponse Response Object
type SetGaussMySqlProxyWeightResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetGaussMySqlProxyWeightResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetGaussMySqlProxyWeightResponse struct{}"
	}

	return strings.Join([]string{"SetGaussMySqlProxyWeightResponse", string(data)}, " ")
}
