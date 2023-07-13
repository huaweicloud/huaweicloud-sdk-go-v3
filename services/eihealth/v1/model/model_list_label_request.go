package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLabelRequest Request Object
type ListLabelRequest struct {
}

func (o ListLabelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLabelRequest struct{}"
	}

	return strings.Join([]string{"ListLabelRequest", string(data)}, " ")
}
