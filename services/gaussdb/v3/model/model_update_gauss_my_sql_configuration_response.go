package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGaussMySqlConfigurationResponse Response Object
type UpdateGaussMySqlConfigurationResponse struct {

	// 修改参数模板的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGaussMySqlConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlConfigurationResponse", string(data)}, " ")
}
