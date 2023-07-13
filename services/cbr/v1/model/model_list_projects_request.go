package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectsRequest Request Object
type ListProjectsRequest struct {
}

func (o ListProjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectsRequest", string(data)}, " ")
}
