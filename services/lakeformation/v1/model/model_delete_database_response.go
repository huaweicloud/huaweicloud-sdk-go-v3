package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatabaseResponse Response Object
type DeleteDatabaseResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseResponse struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseResponse", string(data)}, " ")
}
