package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAuthorizationsRequest struct {
}

func (o ListAuthorizationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthorizationsRequest struct{}"
	}

	return strings.Join([]string{"ListAuthorizationsRequest", string(data)}, " ")
}
