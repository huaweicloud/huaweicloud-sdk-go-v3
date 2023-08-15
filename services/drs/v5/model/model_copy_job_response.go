package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyJobResponse Response Object
type CopyJobResponse struct {

	// 任务ID。
	Id *string `json:"id,omitempty"`

	// 任务名称。
	Name *string `json:"name,omitempty"`

	// 任务状态。
	Status *string `json:"status,omitempty"`

	// 任务创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 是否为克隆任务。
	IsCloneJob     *string `json:"is_clone_job,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopyJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyJobResponse struct{}"
	}

	return strings.Join([]string{"CopyJobResponse", string(data)}, " ")
}
