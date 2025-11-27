package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRecyclePolicyResponse Response Object
type UpdateRecyclePolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateRecyclePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecyclePolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateRecyclePolicyResponse", string(data)}, " ")
}
