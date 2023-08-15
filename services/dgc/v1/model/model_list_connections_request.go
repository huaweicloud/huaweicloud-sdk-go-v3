package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionsRequest Request Object
type ListConnectionsRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`
}

func (o ListConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListConnectionsRequest", string(data)}, " ")
}
