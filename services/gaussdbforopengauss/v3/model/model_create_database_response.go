package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseResponse Response Object
type CreateDatabaseResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseResponse struct{}"
	}

	return strings.Join([]string{"CreateDatabaseResponse", string(data)}, " ")
}
