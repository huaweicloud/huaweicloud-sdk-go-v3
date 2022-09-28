package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateDatabaseDataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDatabaseDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseDataResponse struct{}"
	}

	return strings.Join([]string{"CreateDatabaseDataResponse", string(data)}, " ")
}
