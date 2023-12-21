package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCbhRequest Request Object
type CreateCbhRequest struct {
	Body *CreateInstanceBody `json:"body,omitempty"`
}

func (o CreateCbhRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCbhRequest struct{}"
	}

	return strings.Join([]string{"CreateCbhRequest", string(data)}, " ")
}
