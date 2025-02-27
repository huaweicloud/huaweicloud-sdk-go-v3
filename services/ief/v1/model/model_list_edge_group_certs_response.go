package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEdgeGroupCertsResponse Response Object
type ListEdgeGroupCertsResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListEdgeGroupCertsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeGroupCertsResponse struct{}"
	}

	return strings.Join([]string{"ListEdgeGroupCertsResponse", string(data)}, " ")
}
