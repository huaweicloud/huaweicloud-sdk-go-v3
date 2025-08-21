package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferGroupRequest Request Object
type TransferGroupRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	Body *BussinessGroupTransferBodyDto `json:"body,omitempty"`
}

func (o TransferGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferGroupRequest struct{}"
	}

	return strings.Join([]string{"TransferGroupRequest", string(data)}, " ")
}
