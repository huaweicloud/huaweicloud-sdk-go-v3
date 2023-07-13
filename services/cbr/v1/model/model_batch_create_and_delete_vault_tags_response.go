package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateAndDeleteVaultTagsResponse Response Object
type BatchCreateAndDeleteVaultTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateAndDeleteVaultTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateAndDeleteVaultTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateAndDeleteVaultTagsResponse", string(data)}, " ")
}
