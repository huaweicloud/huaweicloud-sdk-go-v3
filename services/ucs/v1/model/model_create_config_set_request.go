package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConfigSetRequest Request Object
type CreateConfigSetRequest struct {
	Body *CreateConfigSetRequestBody `json:"body,omitempty"`
}

func (o CreateConfigSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigSetRequest struct{}"
	}

	return strings.Join([]string{"CreateConfigSetRequest", string(data)}, " ")
}
