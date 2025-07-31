package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkLoadsResponse Response Object
type ListWorkLoadsResponse struct {
	Body           *[]Workload `json:"body,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListWorkLoadsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkLoadsResponse struct{}"
	}

	return strings.Join([]string{"ListWorkLoadsResponse", string(data)}, " ")
}
