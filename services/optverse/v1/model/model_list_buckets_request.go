package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBucketsRequest Request Object
type ListBucketsRequest struct {
}

func (o ListBucketsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBucketsRequest struct{}"
	}

	return strings.Join([]string{"ListBucketsRequest", string(data)}, " ")
}
