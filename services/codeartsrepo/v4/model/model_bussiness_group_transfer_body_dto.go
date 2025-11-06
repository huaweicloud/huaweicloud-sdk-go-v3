package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BussinessGroupTransferBodyDto 移交代码组
type BussinessGroupTransferBodyDto struct {

	// 移交目标用户id
	OwnerId *int32 `json:"owner_id,omitempty"`
}

func (o BussinessGroupTransferBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BussinessGroupTransferBodyDto struct{}"
	}

	return strings.Join([]string{"BussinessGroupTransferBodyDto", string(data)}, " ")
}
