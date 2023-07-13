package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePersistentStorageResponse Response Object
type DeletePersistentStorageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePersistentStorageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePersistentStorageResponse struct{}"
	}

	return strings.Join([]string{"DeletePersistentStorageResponse", string(data)}, " ")
}
