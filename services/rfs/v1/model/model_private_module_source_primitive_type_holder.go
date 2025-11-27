package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateModuleSourcePrimitiveTypeHolder struct {

	// 在模板中使用模块需要定义如下格式：   module \"my_hello_word_module\" {     source = {module_source}   }  其中{module_source}为本参数
	ModuleSource *string `json:"module_source,omitempty"`
}

func (o PrivateModuleSourcePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateModuleSourcePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateModuleSourcePrimitiveTypeHolder", string(data)}, " ")
}
