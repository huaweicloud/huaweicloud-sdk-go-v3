package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateJobAuthInfoRequest Request Object
type CreateJobAuthInfoRequest struct {
	Body *CreateJobAuthInfoRequestBody `json:"body,omitempty"`
}

func (o CreateJobAuthInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobAuthInfoRequest struct{}"
	}

	return strings.Join([]string{"CreateJobAuthInfoRequest", string(data)}, " ")
}
