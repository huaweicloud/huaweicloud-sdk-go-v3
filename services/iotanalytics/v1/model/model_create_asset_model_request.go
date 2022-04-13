package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
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
