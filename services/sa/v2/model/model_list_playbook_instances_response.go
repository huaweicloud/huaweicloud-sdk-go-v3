package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlaybookInstancesResponse Response Object
type ListPlaybookInstancesResponse struct {

	// tatal count
	Count *int32 `json:"count,omitempty"`

	// list of informations of PlaybookInstanceInfo
	Instances *[]PlaybookInstanceInfo `json:"instances,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPlaybookInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybookInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListPlaybookInstancesResponse", string(data)}, " ")
}
