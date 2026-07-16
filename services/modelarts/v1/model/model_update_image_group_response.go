package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateImageGroupResponse Response Object
type UpdateImageGroupResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateImageGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImageGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateImageGroupResponse", string(data)}, " ")
}
