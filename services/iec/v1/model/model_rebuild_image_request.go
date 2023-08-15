package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebuildImageRequest Request Object
type RebuildImageRequest struct {

	// 租户ID。
	DomainId string `json:"domain_id"`

	// 任务ID。
	JobId string `json:"job_id"`

	Body *RebuildImageRequestBody `json:"body,omitempty"`
}

func (o RebuildImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebuildImageRequest struct{}"
	}

	return strings.Join([]string{"RebuildImageRequest", string(data)}, " ")
}
