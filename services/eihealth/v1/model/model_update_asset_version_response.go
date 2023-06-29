package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssetVersionResponse Response Object
type UpdateAssetVersionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAssetVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssetVersionResponse struct{}"
	}

	return strings.Join([]string{"UpdateAssetVersionResponse", string(data)}, " ")
}
