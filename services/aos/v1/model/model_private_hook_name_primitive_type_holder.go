package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateHookNamePrimitiveTypeHolder struct {

	// 私有hook的名字。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。  推荐用户使用三段命名空间：{自定义hook名称}-{hook应用场景}-hook。
	HookName string `json:"hook_name"`
}

func (o PrivateHookNamePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateHookNamePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateHookNamePrimitiveTypeHolder", string(data)}, " ")
}
