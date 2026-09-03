package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDdsConnectionStatRequest Request Object
type ShowDdsConnectionStatRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	NodeId *string `json:"node_id,omitempty"`

	CurPage *int32 `json:"cur_page,omitempty"`

	PerPage *int32 `json:"per_page,omitempty"`
}

func (o ShowDdsConnectionStatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDdsConnectionStatRequest struct{}"
	}

	return strings.Join([]string{"ShowDdsConnectionStatRequest", string(data)}, " ")
}
