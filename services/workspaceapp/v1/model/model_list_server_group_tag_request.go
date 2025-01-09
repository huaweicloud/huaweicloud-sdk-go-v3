package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerGroupTagRequest Request Object
type ListServerGroupTagRequest struct {
}

func (o ListServerGroupTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerGroupTagRequest struct{}"
	}

	return strings.Join([]string{"ListServerGroupTagRequest", string(data)}, " ")
}
