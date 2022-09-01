package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListConnectionsResponse struct {
	Total *int32 `json:"total,omitempty" xml:"total"`

	Connections    *[]ConnectionInfo `json:"connections,omitempty" xml:"connections"`
	HttpStatusCode int               `json:"-"`
}

func (o ListConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListConnectionsResponse", string(data)}, " ")
}
