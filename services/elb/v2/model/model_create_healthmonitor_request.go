package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateHealthmonitorRequest struct {
	Body *CreateHealthmonitorRequestBody `json:"body,omitempty"`
}

func (o CreateHealthmonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthmonitorRequest struct{}"
	}

	return strings.Join([]string{"CreateHealthmonitorRequest", string(data)}, " ")
}
