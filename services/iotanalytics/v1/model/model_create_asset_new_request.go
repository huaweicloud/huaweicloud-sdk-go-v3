package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAssetNewRequest struct {
	Body *AssetAddRequest `json:"body,omitempty"`
}

func (o CreateAssetNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetNewRequest struct{}"
	}

	return strings.Join([]string{"CreateAssetNewRequest", string(data)}, " ")
}
