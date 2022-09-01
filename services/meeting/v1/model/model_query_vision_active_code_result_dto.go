package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 搜索激活码的返回结果DTO对象
type QueryVisionActiveCodeResultDto struct {

	// 激活码唯一标识
	Id *string `json:"id,omitempty" xml:"id"`

	// 激活码
	ActiveCode *string `json:"activeCode,omitempty" xml:"activeCode"`

	// 终端名称
	DevName *string `json:"devName,omitempty" xml:"devName"`

	// 终端类型
	DevType *string `json:"devType,omitempty" xml:"devType"`

	// 部门编码
	DeptCode *string `json:"deptCode,omitempty" xml:"deptCode"`

	// 部门名称
	DeptName *string `json:"deptName,omitempty" xml:"deptName"`

	// 失效时间戳
	ExpireDate *int64 `json:"expireDate,omitempty" xml:"expireDate"`
}

func (o QueryVisionActiveCodeResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryVisionActiveCodeResultDto struct{}"
	}

	return strings.Join([]string{"QueryVisionActiveCodeResultDto", string(data)}, " ")
}
