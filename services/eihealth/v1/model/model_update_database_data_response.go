package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateDatabaseDataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDatabaseDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseDataResponse struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseDataResponse", string(data)}, " ")
}
