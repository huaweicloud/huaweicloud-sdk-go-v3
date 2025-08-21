package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteRepositoryFilePushPermissionsResponse Response Object
type BatchDeleteRepositoryFilePushPermissionsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteRepositoryFilePushPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRepositoryFilePushPermissionsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteRepositoryFilePushPermissionsResponse", string(data)}, " ")
}
