package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkspacesRequest Request Object
type ListWorkspacesRequest struct {
}

func (o ListWorkspacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkspacesRequest struct{}"
	}

	return strings.Join([]string{"ListWorkspacesRequest", string(data)}, " ")
}
