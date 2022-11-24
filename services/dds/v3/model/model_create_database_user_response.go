package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateDatabaseUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDatabaseUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseUserResponse struct{}"
	}

	return strings.Join([]string{"CreateDatabaseUserResponse", string(data)}, " ")
}
