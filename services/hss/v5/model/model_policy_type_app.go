package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyTypeApp **参数解释**： 进程白名单策略类型 **取值范围**: - allow：允许指定/授权进程运行 - block：阻止潜在恶意软件运行
type PolicyTypeApp struct {
}

func (o PolicyTypeApp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyTypeApp struct{}"
	}

	return strings.Join([]string{"PolicyTypeApp", string(data)}, " ")
}
