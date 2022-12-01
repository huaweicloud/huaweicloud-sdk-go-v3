package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunQueryCustomTagsRequest struct {
}

func (o RunQueryCustomTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunQueryCustomTagsRequest struct{}"
	}

	return strings.Join([]string{"RunQueryCustomTagsRequest", string(data)}, " ")
}
