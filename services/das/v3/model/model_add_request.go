package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddRequest Request Object
type AddRequest struct {
	Body *AddRequestBody `json:"body,omitempty"`
}

func (o AddRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRequest struct{}"
	}

	return strings.Join([]string{"AddRequest", string(data)}, " ")
}
