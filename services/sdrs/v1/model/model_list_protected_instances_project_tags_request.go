package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListProtectedInstancesProjectTagsRequest struct {
}

func (o ListProtectedInstancesProjectTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectedInstancesProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ListProtectedInstancesProjectTagsRequest", string(data)}, " ")
}
