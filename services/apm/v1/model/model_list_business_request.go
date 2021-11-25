package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBusinessRequest struct {
}

func (o ListBusinessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBusinessRequest struct{}"
	}

	return strings.Join([]string{"ListBusinessRequest", string(data)}, " ")
}
