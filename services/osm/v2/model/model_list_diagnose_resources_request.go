package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDiagnoseResourcesRequest Request Object
type ListDiagnoseResourcesRequest struct {
	Body *QueryTscDiagnoseResourcesReq `json:"body,omitempty"`
}

func (o ListDiagnoseResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiagnoseResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListDiagnoseResourcesRequest", string(data)}, " ")
}
