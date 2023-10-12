package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSessionTypeRequest Request Object
type ListSessionTypeRequest struct {
}

func (o ListSessionTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSessionTypeRequest struct{}"
	}

	return strings.Join([]string{"ListSessionTypeRequest", string(data)}, " ")
}
