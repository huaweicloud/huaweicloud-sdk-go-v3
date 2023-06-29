package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFuncSnapshotRequest Request Object
type UpdateFuncSnapshotRequest struct {

	// 禁用/启用
	Action string `json:"action"`

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`
}

func (o UpdateFuncSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFuncSnapshotRequest struct{}"
	}

	return strings.Join([]string{"UpdateFuncSnapshotRequest", string(data)}, " ")
}
