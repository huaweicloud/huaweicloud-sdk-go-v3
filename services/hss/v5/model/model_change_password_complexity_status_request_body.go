package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePasswordComplexityStatusRequestBody 请求参数
type ChangePasswordComplexityStatusRequestBody struct {

	// 是否是全量操作。每次最多处理1000个主机。
	OperateAll *bool `json:"operate_all,omitempty"`

	// 主机id列表。operate_all=ture时不处理host_ids参数。
	HostIds *[]string `json:"host_ids,omitempty"`
}

func (o ChangePasswordComplexityStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePasswordComplexityStatusRequestBody struct{}"
	}

	return strings.Join([]string{"ChangePasswordComplexityStatusRequestBody", string(data)}, " ")
}
