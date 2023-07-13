package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssetModelRequest Request Object
type CreateAssetModelRequest struct {
	Body *AssetModelAddRequest `json:"body,omitempty"`
}

func (o CreateAssetModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetModelRequest struct{}"
	}

	return strings.Join([]string{"CreateAssetModelRequest", string(data)}, " ")
}
