package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkspacesResponse Response Object
type ListWorkspacesResponse struct {
	Data           *ListWorkspacesResultData `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListWorkspacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkspacesResponse struct{}"
	}

	return strings.Join([]string{"ListWorkspacesResponse", string(data)}, " ")
}
