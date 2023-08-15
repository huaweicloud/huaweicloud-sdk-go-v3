package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrUpdateAssetResponse Response Object
type CreateOrUpdateAssetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateOrUpdateAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateAssetResponse struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateAssetResponse", string(data)}, " ")
}
