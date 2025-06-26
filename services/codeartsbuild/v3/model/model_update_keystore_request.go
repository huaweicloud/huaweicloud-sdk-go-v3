package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKeystoreRequest Request Object
type UpdateKeystoreRequest struct {

	// 文件ID
	Id string `json:"id"`

	Body *UpdateKeystoreRequestBody `json:"body,omitempty"`
}

func (o UpdateKeystoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeystoreRequest struct{}"
	}

	return strings.Join([]string{"UpdateKeystoreRequest", string(data)}, " ")
}
