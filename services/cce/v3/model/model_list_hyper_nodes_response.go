package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHyperNodesResponse Response Object
type ListHyperNodesResponse struct {
	Body           *[]HyperNode `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListHyperNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHyperNodesResponse struct{}"
	}

	return strings.Join([]string{"ListHyperNodesResponse", string(data)}, " ")
}
