package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProtectedInstanceNameResponse Response Object
type UpdateProtectedInstanceNameResponse struct {
	ProtectedInstance *ShowProtectedInstanceParams `json:"protected_instance,omitempty"`
	HttpStatusCode    int                          `json:"-"`
}

func (o UpdateProtectedInstanceNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProtectedInstanceNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateProtectedInstanceNameResponse", string(data)}, " ")
}
