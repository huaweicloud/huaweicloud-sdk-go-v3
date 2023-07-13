package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatabaseDataResponse Response Object
type DeleteDatabaseDataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDatabaseDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseDataResponse struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseDataResponse", string(data)}, " ")
}
