package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyConfigurationWithEntitiesResponse Response Object
type ApplyConfigurationWithEntitiesResponse struct {

	// 应用参数模板的异步任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ApplyConfigurationWithEntitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyConfigurationWithEntitiesResponse struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationWithEntitiesResponse", string(data)}, " ")
}
