package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateObjectConfigDescRequest Request Object
type UpdateObjectConfigDescRequest struct {
	Body *UpdateObjectConfigDesc `json:"body,omitempty"`
}

func (o UpdateObjectConfigDescRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateObjectConfigDescRequest struct{}"
	}

	return strings.Join([]string{"UpdateObjectConfigDescRequest", string(data)}, " ")
}
