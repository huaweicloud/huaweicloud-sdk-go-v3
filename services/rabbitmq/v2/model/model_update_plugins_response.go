package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePluginsResponse Response Object
type UpdatePluginsResponse struct {

	// 后台任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePluginsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePluginsResponse struct{}"
	}

	return strings.Join([]string{"UpdatePluginsResponse", string(data)}, " ")
}
