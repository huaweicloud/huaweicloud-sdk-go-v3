package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type EnableLogCollectionRequest struct {
}

func (o EnableLogCollectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableLogCollectionRequest struct{}"
	}

	return strings.Join([]string{"EnableLogCollectionRequest", string(data)}, " ")
}
