package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteDatabaseUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDatabaseUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseUserResponse", string(data)}, " ")
}
