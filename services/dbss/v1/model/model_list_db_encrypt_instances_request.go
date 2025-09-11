package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbEncryptInstancesRequest Request Object
type ListDbEncryptInstancesRequest struct {
}

func (o ListDbEncryptInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbEncryptInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListDbEncryptInstancesRequest", string(data)}, " ")
}
