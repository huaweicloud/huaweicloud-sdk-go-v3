package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDiagnoseItemsResponse Response Object
type ListDiagnoseItemsResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 检查项结果列表
	ItemResults    *[]ItemResultDetailVo `json:"item_results,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListDiagnoseItemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiagnoseItemsResponse struct{}"
	}

	return strings.Join([]string{"ListDiagnoseItemsResponse", string(data)}, " ")
}
