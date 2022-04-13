package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateWhitelistRequest struct {
	Body *CreateWhitelistRequestBody `json:"body,omitempty"`
}

func (o CreateWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWhitelistRequest struct{}"
	}

	return strings.Join([]string{"CreateWhitelistRequest", string(data)}, " ")
}
