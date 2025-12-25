package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNetProxyRequest Request Object
type ListNetProxyRequest struct {
}

func (o ListNetProxyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNetProxyRequest struct{}"
	}

	return strings.Join([]string{"ListNetProxyRequest", string(data)}, " ")
}
