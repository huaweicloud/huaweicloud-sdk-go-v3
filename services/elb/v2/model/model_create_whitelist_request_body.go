package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWhitelistRequestBody This is a auto create Body Object
type CreateWhitelistRequestBody struct {
	Whitelist *CreateWhitelistReq `json:"whitelist"`
}

func (o CreateWhitelistRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWhitelistRequestBody struct{}"
	}

	return strings.Join([]string{"CreateWhitelistRequestBody", string(data)}, " ")
}
