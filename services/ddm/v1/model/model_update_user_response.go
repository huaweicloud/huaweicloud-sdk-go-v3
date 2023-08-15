package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserResponse Response Object
type UpdateUserResponse struct {

	// DDM实例帐号名称。
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserResponse", string(data)}, " ")
}
