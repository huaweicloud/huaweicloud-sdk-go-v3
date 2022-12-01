package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunDeleteCustomTagsRequest struct {
}

func (o RunDeleteCustomTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDeleteCustomTagsRequest struct{}"
	}

	return strings.Join([]string{"RunDeleteCustomTagsRequest", string(data)}, " ")
}
