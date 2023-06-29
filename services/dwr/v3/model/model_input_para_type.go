package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InputParaType Input结构体参数类型。支持string,integer,float,boolean,list,map
type InputParaType struct {
}

func (o InputParaType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InputParaType struct{}"
	}

	return strings.Join([]string{"InputParaType", string(data)}, " ")
}
