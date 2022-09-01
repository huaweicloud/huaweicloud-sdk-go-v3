package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateWhitelistRequest struct {
	Body *CreateWhitelistRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWhitelistRequest struct{}"
	}

	return strings.Join([]string{"CreateWhitelistRequest", string(data)}, " ")
}
