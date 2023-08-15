package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeEnvironmentRequest Request Object
type ChangeEnvironmentRequest struct {

	// 环境ID。
	EnvironmentId string `json:"environment_id"`

	Body *EnvironmentModify `json:"body,omitempty"`
}

func (o ChangeEnvironmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeEnvironmentRequest struct{}"
	}

	return strings.Join([]string{"ChangeEnvironmentRequest", string(data)}, " ")
}
