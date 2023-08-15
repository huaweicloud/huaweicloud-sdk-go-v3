package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessPolicyObjectsResponse Response Object
type UpdateAccessPolicyObjectsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAccessPolicyObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessPolicyObjectsResponse struct{}"
	}

	return strings.Join([]string{"UpdateAccessPolicyObjectsResponse", string(data)}, " ")
}
