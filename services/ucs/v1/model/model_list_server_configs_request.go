package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerConfigsRequest Request Object
type ListServerConfigsRequest struct {
}

func (o ListServerConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListServerConfigsRequest", string(data)}, " ")
}
