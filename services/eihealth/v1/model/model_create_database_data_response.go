package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseDataResponse Response Object
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
