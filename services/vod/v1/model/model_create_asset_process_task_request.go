package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAssetProcessTaskRequest struct {
	Body *AssetProcessReq `json:"body,omitempty" xml:"body"`
}

func (o CreateAssetProcessTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetProcessTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateAssetProcessTaskRequest", string(data)}, " ")
}
