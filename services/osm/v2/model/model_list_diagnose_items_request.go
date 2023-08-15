package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDiagnoseItemsRequest Request Object
type ListDiagnoseItemsRequest struct {
	Body *QueryDiagnoseItemsReq `json:"body,omitempty"`
}

func (o ListDiagnoseItemsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiagnoseItemsRequest struct{}"
	}

	return strings.Join([]string{"ListDiagnoseItemsRequest", string(data)}, " ")
}
