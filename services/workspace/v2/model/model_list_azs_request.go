package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAzsRequest Request Object
type ListAzsRequest struct {
}

func (o ListAzsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAzsRequest struct{}"
	}

	return strings.Join([]string{"ListAzsRequest", string(data)}, " ")
}
