package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOperationResponse Response Object
type ListOperationResponse struct {

	// 操作ID。
	OperationId *string `json:"operation_id,omitempty"`

	// 完成进度百分比。
	PercentageComplete *int32 `json:"percentage_complete,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 创建账号、纳管注册OU、纳管账号的详细进度信息。
	PercentageDetails *[]OrganizationalPercentageDetail `json:"percentage_details,omitempty"`

	// 创建账号、纳管注册OU、纳管账号的错误信息描述。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListOperationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOperationResponse struct{}"
	}

	return strings.Join([]string{"ListOperationResponse", string(data)}, " ")
}
