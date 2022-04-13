package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAssetProcessTaskRequest struct {
	Body *AssetProcessReq `json:"body,omitempty"`
}

func (o CreateAssetProcessTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetProcessTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateAssetProcessTaskRequest", string(data)}, " ")
}
