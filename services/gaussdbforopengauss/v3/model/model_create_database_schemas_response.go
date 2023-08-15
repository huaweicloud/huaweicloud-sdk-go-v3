package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseSchemasResponse Response Object
type CreateDatabaseSchemasResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDatabaseSchemasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseSchemasResponse struct{}"
	}

	return strings.Join([]string{"CreateDatabaseSchemasResponse", string(data)}, " ")
}
