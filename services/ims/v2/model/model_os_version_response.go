package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询版本响应体
type OsVersionResponse struct {

	// 接口状态。
	Status string `json:"status"`

	// 接口ID。
	Id *string `json:"id,omitempty"`

	// 自描述信息。
	Links *[]Links `json:"links,omitempty"`
}

func (o OsVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsVersionResponse struct{}"
	}

	return strings.Join([]string{"OsVersionResponse", string(data)}, " ")
}
