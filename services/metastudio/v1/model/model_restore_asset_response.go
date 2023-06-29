package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreAssetResponse Response Object
type RestoreAssetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestoreAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreAssetResponse struct{}"
	}

	return strings.Join([]string{"RestoreAssetResponse", string(data)}, " ")
}
