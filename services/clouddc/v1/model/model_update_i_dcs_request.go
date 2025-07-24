package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIDcsRequest Request Object
type UpdateIDcsRequest struct {
	Body *[]interface{} `json:"body,omitempty"`
}

func (o UpdateIDcsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIDcsRequest struct{}"
	}

	return strings.Join([]string{"UpdateIDcsRequest", string(data)}, " ")
}
