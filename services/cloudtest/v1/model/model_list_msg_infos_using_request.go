package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMsgInfosUsingRequest Request Object
type ListMsgInfosUsingRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	Body *MsgInfoQuery `json:"body,omitempty"`
}

func (o ListMsgInfosUsingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMsgInfosUsingRequest struct{}"
	}

	return strings.Join([]string{"ListMsgInfosUsingRequest", string(data)}, " ")
}
