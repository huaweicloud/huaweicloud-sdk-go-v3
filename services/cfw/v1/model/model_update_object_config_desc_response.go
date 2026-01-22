package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateObjectConfigDescResponse Response Object
type UpdateObjectConfigDescResponse struct {

	// 成员id
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateObjectConfigDescResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateObjectConfigDescResponse struct{}"
	}

	return strings.Join([]string{"UpdateObjectConfigDescResponse", string(data)}, " ")
}
