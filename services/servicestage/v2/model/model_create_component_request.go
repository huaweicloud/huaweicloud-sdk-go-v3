package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateComponentRequest Request Object
type CreateComponentRequest struct {

	// 应用ID。
	ApplicationId string `json:"application_id"`

	Body *ComponentCreate `json:"body,omitempty"`
}

func (o CreateComponentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentRequest struct{}"
	}

	return strings.Join([]string{"CreateComponentRequest", string(data)}, " ")
}
