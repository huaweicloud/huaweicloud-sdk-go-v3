package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGrantResponse Response Object
type UpdateGrantResponse struct {
	Data           *GrantData `json:"data,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o UpdateGrantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGrantResponse struct{}"
	}

	return strings.Join([]string{"UpdateGrantResponse", string(data)}, " ")
}
