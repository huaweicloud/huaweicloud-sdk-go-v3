package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseUserResponse Response Object
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
