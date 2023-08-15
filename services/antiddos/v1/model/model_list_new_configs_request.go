package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNewConfigsRequest Request Object
type ListNewConfigsRequest struct {
}

func (o ListNewConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNewConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListNewConfigsRequest", string(data)}, " ")
}
