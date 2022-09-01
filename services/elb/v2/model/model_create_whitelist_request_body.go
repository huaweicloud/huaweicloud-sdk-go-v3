package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CreateWhitelistRequestBody struct {
	Whitelist *CreateWhitelistReq `json:"whitelist" xml:"whitelist"`
}

func (o CreateWhitelistRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWhitelistRequestBody struct{}"
	}

	return strings.Join([]string{"CreateWhitelistRequestBody", string(data)}, " ")
}
