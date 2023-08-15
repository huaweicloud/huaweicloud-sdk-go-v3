package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionsResponse Response Object
type ListConnectionsResponse struct {
	Total *int32 `json:"total,omitempty"`

	Connections    *[]ConnectionInfo `json:"connections,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListConnectionsResponse", string(data)}, " ")
}
