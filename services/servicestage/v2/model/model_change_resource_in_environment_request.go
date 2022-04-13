package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeResourceInEnvironmentRequest struct {
	// 环境ID。

	EnvironmentId string `json:"environment_id"`

	Body *EnvironmentResourceModify `json:"body,omitempty"`
}

func (o ChangeResourceInEnvironmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeResourceInEnvironmentRequest struct{}"
	}

	return strings.Join([]string{"ChangeResourceInEnvironmentRequest", string(data)}, " ")
}
