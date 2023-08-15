package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreOldInstanceRequest Request Object
type RestoreOldInstanceRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *RestoreRequest `json:"body,omitempty"`
}

func (o RestoreOldInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreOldInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreOldInstanceRequest", string(data)}, " ")
}
