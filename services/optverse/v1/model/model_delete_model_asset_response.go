package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteModelAssetResponse Response Object
type DeleteModelAssetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteModelAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteModelAssetResponse struct{}"
	}

	return strings.Join([]string{"DeleteModelAssetResponse", string(data)}, " ")
}
