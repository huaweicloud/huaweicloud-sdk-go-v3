package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckHealthRequestBody struct {

	// version
	Version *int32 `bson:"version,omitempty"`
}

func (o CheckHealthRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckHealthRequestBody struct{}"
	}

	return strings.Join([]string{"CheckHealthRequestBody", string(data)}, " ")
}
