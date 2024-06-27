package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateVersionTestCasesResponse Response Object
type BatchUpdateVersionTestCasesResponse struct {

	// CTS需要返回资源id
	Id *string `json:"id,omitempty"`

	// CTS需要返回资源name
	Name *string `json:"name,omitempty"`

	// 成功批量更新用例的id列表
	SuccessList *[]string `json:"success_list,omitempty"`

	// 没有批量更新用例的id列表
	FailedList     *[]string `json:"failed_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchUpdateVersionTestCasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateVersionTestCasesResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateVersionTestCasesResponse", string(data)}, " ")
}
