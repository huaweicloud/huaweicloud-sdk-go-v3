package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventOperatesResponse Response Object
type ListEventOperatesResponse struct {

	// 支持的处理操作
	OperateAcceptList *[]string `json:"operate_accept_list,omitempty"`
	HttpStatusCode    int       `json:"-"`
}

func (o ListEventOperatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventOperatesResponse struct{}"
	}

	return strings.Join([]string{"ListEventOperatesResponse", string(data)}, " ")
}
