package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListArchiveConfigsRequest Request Object
type ListArchiveConfigsRequest struct {
}

func (o ListArchiveConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListArchiveConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListArchiveConfigsRequest", string(data)}, " ")
}
