package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DecoyPortBasePortStruct 诱饵端口基础结构
type DecoyPortBasePortStruct struct {

	// **参数解释**： 端口号 **取值范围**： 最小值0，最大值65535
	Port *int32 `json:"port,omitempty"`

	// **参数解释**： 协议类型 **取值范围**： - tcp - tcp6
	Protocol *string `json:"protocol,omitempty"`
}

func (o DecoyPortBasePortStruct) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecoyPortBasePortStruct struct{}"
	}

	return strings.Join([]string{"DecoyPortBasePortStruct", string(data)}, " ")
}
