package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllJobByNameRequest Request Object
type ListAllJobByNameRequest struct {
	Body *SearchJobsRequestBody `json:"body,omitempty"`
}

func (o ListAllJobByNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllJobByNameRequest struct{}"
	}

	return strings.Join([]string{"ListAllJobByNameRequest", string(data)}, " ")
}
