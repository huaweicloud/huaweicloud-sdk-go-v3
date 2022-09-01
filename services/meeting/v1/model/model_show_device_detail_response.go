package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDeviceDetailResponse struct {

	// 终端名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 终端类型，区分自研和第三方终端。
	Type *string `json:"type,omitempty" xml:"type"`

	// 终端型号，枚举类型。当前支持TE系列硬件终端，具体的终端类型可以通过获取所有终端类型接口查询。
	Model *string `json:"model,omitempty" xml:"model"`

	// 终端SN号，仅可包含数字、字母和下划线。
	Sn *string `json:"sn,omitempty" xml:"sn"`

	// 硬终端对应的内置账号。
	Account *string `json:"account,omitempty" xml:"account"`

	// 终端绑定的号码
	Number *string `json:"number,omitempty" xml:"number"`

	// 投影码生成模式 * 0、自动(该模式下根据消息上报的IP地址内部控制复杂度：私网地址配置成简单模式；公网地址配置成复杂模式) * 1、简单 * 2、复杂
	PrjCodeMode *int32 `json:"prjCodeMode,omitempty" xml:"prjCodeMode"`

	// 部门编号
	DeptCode *string `json:"deptCode,omitempty" xml:"deptCode"`

	// 部门名称
	DeptName *string `json:"deptName,omitempty" xml:"deptName"`

	// 部门名称路径
	DeptNamePath *string `json:"deptNamePath,omitempty" xml:"deptNamePath"`

	// 手机号
	Phone *string `json:"phone,omitempty" xml:"phone"`

	// 手机号所属的国家
	Country *string `json:"country,omitempty" xml:"country"`

	// 邮箱
	Email *string `json:"email,omitempty" xml:"email"`

	// 终端描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 终端状态 * 0、正常 * 1、停用\"
	Status         *int32 `json:"status,omitempty" xml:"status"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDeviceDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowDeviceDetailResponse", string(data)}, " ")
}
