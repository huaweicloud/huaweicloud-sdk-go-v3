package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyWorkspaceResponse Response Object
type ApplyWorkspaceResponse struct {

	// 开通云办公服务的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ApplyWorkspaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyWorkspaceResponse struct{}"
	}

	return strings.Join([]string{"ApplyWorkspaceResponse", string(data)}, " ")
}
