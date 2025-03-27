package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConsistencyResultResponse Response Object
type UpdateConsistencyResultResponse struct {

	// 更新一致性校验结果成功。
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateConsistencyResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConsistencyResultResponse struct{}"
	}

	return strings.Join([]string{"UpdateConsistencyResultResponse", string(data)}, " ")
}
