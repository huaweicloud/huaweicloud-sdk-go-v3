package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateObjectRetrievalResponse Response Object
type CreateObjectRetrievalResponse struct {

	// 取回数据的任务id
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateObjectRetrievalResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateObjectRetrievalResponse struct{}"
	}

	return strings.Join([]string{"CreateObjectRetrievalResponse", string(data)}, " ")
}
