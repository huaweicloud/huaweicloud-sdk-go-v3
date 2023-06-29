package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWhitelistRequestBody This is a auto create Body Object
type UpdateWhitelistRequestBody struct {
	Whitelist *UpdateWhitelistReq `json:"whitelist"`
}

func (o UpdateWhitelistRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWhitelistRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateWhitelistRequestBody", string(data)}, " ")
}
