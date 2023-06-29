package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseResourceResponse Response Object
type CreateDatabaseResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDatabaseResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseResourceResponse struct{}"
	}

	return strings.Join([]string{"CreateDatabaseResourceResponse", string(data)}, " ")
}
