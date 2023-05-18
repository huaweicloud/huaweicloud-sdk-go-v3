package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPipelineSimpleInfoRequest struct {
	Body *ListPipelineSimpleInfoRequestBody `json:"body,omitempty"`
}

func (o ListPipelineSimpleInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineSimpleInfoRequest struct{}"
	}

	return strings.Join([]string{"ListPipelineSimpleInfoRequest", string(data)}, " ")
}
