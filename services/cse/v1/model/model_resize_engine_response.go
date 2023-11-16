package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeEngineResponse Response Object
type ResizeEngineResponse struct {

	// 引擎id
	Id *string `json:"id,omitempty"`

	// 引擎名字
	Name *string `json:"name,omitempty"`

	// jobID
	JobId          *int32 `json:"job_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ResizeEngineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeEngineResponse struct{}"
	}

	return strings.Join([]string{"ResizeEngineResponse", string(data)}, " ")
}
