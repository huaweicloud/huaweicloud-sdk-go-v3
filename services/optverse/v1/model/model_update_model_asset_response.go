package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateModelAssetResponse Response Object
type UpdateModelAssetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateModelAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModelAssetResponse struct{}"
	}

	return strings.Join([]string{"UpdateModelAssetResponse", string(data)}, " ")
}
