package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteDataStoreResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDataStoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataStoreResponse struct{}"
	}

	return strings.Join([]string{"DeleteDataStoreResponse", string(data)}, " ")
}
