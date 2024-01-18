package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeMasterStandbyAsyncResponse Response Object
type ChangeMasterStandbyAsyncResponse struct {

	// 主备切换任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeMasterStandbyAsyncResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeMasterStandbyAsyncResponse struct{}"
	}

	return strings.Join([]string{"ChangeMasterStandbyAsyncResponse", string(data)}, " ")
}
