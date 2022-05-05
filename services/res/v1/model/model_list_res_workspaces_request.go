package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListResWorkspacesRequest struct {
}

func (o ListResWorkspacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResWorkspacesRequest struct{}"
	}

	return strings.Join([]string{"ListResWorkspacesRequest", string(data)}, " ")
}
