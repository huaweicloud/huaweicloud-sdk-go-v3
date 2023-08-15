package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetAdministratorResponse Response Object
type ResetAdministratorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetAdministratorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetAdministratorResponse struct{}"
	}

	return strings.Join([]string{"ResetAdministratorResponse", string(data)}, " ")
}
