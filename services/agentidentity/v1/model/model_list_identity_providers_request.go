package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIdentityProvidersRequest Request Object
type ListIdentityProvidersRequest struct {
}

func (o ListIdentityProvidersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIdentityProvidersRequest struct{}"
	}

	return strings.Join([]string{"ListIdentityProvidersRequest", string(data)}, " ")
}
