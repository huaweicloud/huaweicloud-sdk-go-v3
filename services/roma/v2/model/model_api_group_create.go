package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiGroupCreate struct {
	// API分组的名称。  支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头，3-255个字符。 > 中文字符必须为UTF-8或者unicode编码。

	Name string `json:"name"`
	// API分组描述。 > 中文字符必须为UTF-8或者unicode编码。

	Remark *string `json:"remark,omitempty"`
	// 分组归属的集成应用编号。  分组版本V2时必填。

	RomaAppId *string `json:"roma_app_id,omitempty"`
	// 分组版本  - V1：全局分组 - V2：应用级分组

	Version *string `json:"version,omitempty"`
}

func (o ApiGroupCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiGroupCreate struct{}"
	}

	return strings.Join([]string{"ApiGroupCreate", string(data)}, " ")
}
