package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssetProcessTaskResponse Response Object
type CreateAssetProcessTaskResponse struct {

	// 媒资ID。
	AssetId        *string `json:"asset_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAssetProcessTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetProcessTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateAssetProcessTaskResponse", string(data)}, " ")
}
