package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAzRequest Request Object
type ListAzRequest struct {
}

func (o ListAzRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAzRequest struct{}"
	}

	return strings.Join([]string{"ListAzRequest", string(data)}, " ")
}
