package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnableAutoRollbackPrimitiveTypeHolder struct {

	// 自动回滚的标识位，如果不传默认为false，即默认不开启资源栈自动回滚（自动回滚开启后，如果部署失败，则会自动回滚，并返回上一个稳定状态）  *在UpdateStack API中，如果该参数未在RequestBody中给予，则不会对资源栈的自动回滚属性进行更新* *该属性与使用模板导入资源功能互斥，如果堆栈的自动回滚设置为true，则不允许部署包含导入资源的模板*
	EnableAutoRollback *bool `json:"enable_auto_rollback,omitempty"`
}

func (o EnableAutoRollbackPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableAutoRollbackPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"EnableAutoRollbackPrimitiveTypeHolder", string(data)}, " ")
}
