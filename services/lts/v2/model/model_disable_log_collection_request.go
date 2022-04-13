package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisableLogCollectionRequest struct {
}

func (o DisableLogCollectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableLogCollectionRequest struct{}"
	}

	return strings.Join([]string{"DisableLogCollectionRequest", string(data)}, " ")
}
