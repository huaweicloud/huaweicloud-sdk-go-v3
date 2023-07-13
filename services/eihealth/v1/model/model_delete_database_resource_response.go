package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatabaseResourceResponse Response Object
type DeleteDatabaseResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDatabaseResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseResourceResponse struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseResourceResponse", string(data)}, " ")
}
