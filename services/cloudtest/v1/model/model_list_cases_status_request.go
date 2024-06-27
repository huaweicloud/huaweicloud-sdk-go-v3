package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCasesStatusRequest Request Object
type ListCasesStatusRequest struct {

	// projectUUId
	TestServiceId string `json:"testServiceId"`

	Body *string `json:"body,omitempty"`
}

func (o ListCasesStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCasesStatusRequest struct{}"
	}

	return strings.Join([]string{"ListCasesStatusRequest", string(data)}, " ")
}
