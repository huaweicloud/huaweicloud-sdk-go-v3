package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
