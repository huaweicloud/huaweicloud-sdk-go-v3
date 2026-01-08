package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteListenersResponse Response Object
type BatchDeleteListenersResponse struct {

	// 批量删除任务id
	JobId *string `json:"job_id,omitempty"`

	// 待删除的监听器id列表。
	ListenerIds    *[]string `json:"listener_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteListenersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteListenersResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteListenersResponse", string(data)}, " ")
}
