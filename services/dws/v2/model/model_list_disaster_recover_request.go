package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDisasterRecoverRequest Request Object
type ListDisasterRecoverRequest struct {
}

func (o ListDisasterRecoverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDisasterRecoverRequest struct{}"
	}

	return strings.Join([]string{"ListDisasterRecoverRequest", string(data)}, " ")
}
