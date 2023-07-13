package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConnectionResponse Response Object
type UpdateConnectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConnectionResponse struct{}"
	}

	return strings.Join([]string{"UpdateConnectionResponse", string(data)}, " ")
}
