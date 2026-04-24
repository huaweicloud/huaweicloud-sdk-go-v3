package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompareUserInfo struct {

	// id
	Id string `json:"id"`

	// 源数据库账号名称
	SrcUserName string `json:"src_user_name"`

	// 目标数据库账号名称
	TarUserName string `json:"tar_user_name"`

	// 源数据库账号权限
	SrcPrivileges string `json:"src_privileges"`

	// 目标数据库账号权限
	TarPrivileges string `json:"tar_privileges"`

	// 目标端是否存在，取值： - false：不存在 - true：存在
	IsTargetExisted bool `json:"is_target_existed"`

	// 对比结果，取值： - INCONSISTENT：不一致 - UNABLE_TO_COMPARE：无法比对 - CONSISTENT：一致 - TARGET_SCHEMA_NOT_EXIST：目标库不存在 - COMPARE_FAILED：比对失败 - COMPARING：比对中 - WAITING_COMPARE：等待比对 - UNKNOWN：未知错误
	CompareResult int32 `json:"compare_result"`

	// 创建时间
	CreatedAt string `json:"created_at"`

	// 更新时间
	UpdatedAt string `json:"updated_at"`
}

func (o CompareUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareUserInfo struct{}"
	}

	return strings.Join([]string{"CompareUserInfo", string(data)}, " ")
}
