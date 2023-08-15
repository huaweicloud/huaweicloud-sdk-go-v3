package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreInstanceFromCollectionResponse Response Object
type RestoreInstanceFromCollectionResponse struct {

	// 库表级恢复的异步任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreInstanceFromCollectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInstanceFromCollectionResponse struct{}"
	}

	return strings.Join([]string{"RestoreInstanceFromCollectionResponse", string(data)}, " ")
}
