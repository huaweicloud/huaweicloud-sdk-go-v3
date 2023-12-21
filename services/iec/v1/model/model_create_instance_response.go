package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceResponse Response Object
type CreateInstanceResponse struct {
	JobId *string `json:"job_id,omitempty"`

	// 创建中实例的ID。
	ServerIds      *[]string `json:"server_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceResponse", string(data)}, " ")
}
