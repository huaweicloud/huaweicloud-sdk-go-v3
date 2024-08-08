package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStoragePolicyStatementResponse Response Object
type ListStoragePolicyStatementResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 支持的访问策略。
	Items          *[]PolicyStatement `json:"items,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListStoragePolicyStatementResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStoragePolicyStatementResponse struct{}"
	}

	return strings.Join([]string{"ListStoragePolicyStatementResponse", string(data)}, " ")
}
