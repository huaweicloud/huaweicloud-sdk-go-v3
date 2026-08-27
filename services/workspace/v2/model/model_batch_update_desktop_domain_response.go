package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateDesktopDomainResponse Response Object
type BatchUpdateDesktopDomainResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchUpdateDesktopDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateDesktopDomainResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateDesktopDomainResponse", string(data)}, " ")
}
