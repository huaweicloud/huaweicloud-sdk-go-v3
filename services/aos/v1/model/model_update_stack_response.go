package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStackResponse Response Object
type UpdateStackResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateStackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStackResponse struct{}"
	}

	return strings.Join([]string{"UpdateStackResponse", string(data)}, " ")
}
