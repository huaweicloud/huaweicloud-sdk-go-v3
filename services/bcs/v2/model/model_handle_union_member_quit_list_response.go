package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type HandleUnionMemberQuitListResponse struct {

	// 请求结果
	Result *string `json:"result,omitempty" xml:"result"`

	// 操作记录ID
	OperationId    *string `json:"operation_id,omitempty" xml:"operation_id"`
	HttpStatusCode int     `json:"-"`
}

func (o HandleUnionMemberQuitListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleUnionMemberQuitListResponse struct{}"
	}

	return strings.Join([]string{"HandleUnionMemberQuitListResponse", string(data)}, " ")
}
