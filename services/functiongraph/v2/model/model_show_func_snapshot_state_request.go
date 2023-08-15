package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFuncSnapshotStateRequest Request Object
type ShowFuncSnapshotStateRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`
}

func (o ShowFuncSnapshotStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFuncSnapshotStateRequest struct{}"
	}

	return strings.Join([]string{"ShowFuncSnapshotStateRequest", string(data)}, " ")
}
