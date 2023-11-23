package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFunctionCollectStateRequest Request Object
type UpdateFunctionCollectStateRequest struct {

	// 函数URN
	FuncUrn string `json:"func_urn"`

	// 置顶状态
	State string `json:"state"`
}

func (o UpdateFunctionCollectStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionCollectStateRequest struct{}"
	}

	return strings.Join([]string{"UpdateFunctionCollectStateRequest", string(data)}, " ")
}
