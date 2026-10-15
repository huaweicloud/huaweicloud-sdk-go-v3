package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportJobListResponse Response Object
type ExportJobListResponse struct {

	// 异步任务ID。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportJobListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportJobListResponse struct{}"
	}

	return strings.Join([]string{"ExportJobListResponse", string(data)}, " ")
}
