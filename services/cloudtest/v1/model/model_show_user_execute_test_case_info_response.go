package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserExecuteTestCaseInfoResponse Response Object
type ShowUserExecuteTestCaseInfoResponse struct {

	// 起始记录数大于实际总条数时，值为0。
	Total *int32 `json:"total,omitempty"`

	// 时段内用例的执行情况
	Values         *[]ExternalUserExecuteInfo `json:"values,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowUserExecuteTestCaseInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserExecuteTestCaseInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowUserExecuteTestCaseInfoResponse", string(data)}, " ")
}
