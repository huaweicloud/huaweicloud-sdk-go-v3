package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImChannelsRequest Request Object
type ListImChannelsRequest struct {

	// Agent 实例主键 ID
	Id string `json:"id"`
}

func (o ListImChannelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImChannelsRequest struct{}"
	}

	return strings.Join([]string{"ListImChannelsRequest", string(data)}, " ")
}
