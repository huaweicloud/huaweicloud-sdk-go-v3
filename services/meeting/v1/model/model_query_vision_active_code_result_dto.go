package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 激活码信息。
type QueryVisionActiveCodeResultDto struct {

	// 激活码唯一标识。
	Id *string `json:"id,omitempty"`

	// 激活码。
	ActiveCode *string `json:"activeCode,omitempty"`

	// 终端名称。
	DevName *string `json:"devName,omitempty"`

	// 终端类型。
	DevType *string `json:"devType,omitempty"`

	// 部门编码。
	DeptCode *string `json:"deptCode,omitempty"`

	// 部门名称。
	DeptName *string `json:"deptName,omitempty"`

	// 失效时间戳。
	ExpireDate *int64 `json:"expireDate,omitempty"`
}

func (o QueryVisionActiveCodeResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryVisionActiveCodeResultDto struct{}"
	}

	return strings.Join([]string{"QueryVisionActiveCodeResultDto", string(data)}, " ")
}
