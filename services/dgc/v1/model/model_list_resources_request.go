package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesRequest Request Object
type ListResourcesRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`
}

func (o ListResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesRequest", string(data)}, " ")
}
