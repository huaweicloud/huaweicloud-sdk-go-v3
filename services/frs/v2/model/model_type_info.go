package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TypeInfo struct {

	// 数据类型:string,integer,double,long等。
	Type *string `json:"type,omitempty"`
}

func (o TypeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TypeInfo struct{}"
	}

	return strings.Join([]string{"TypeInfo", string(data)}, " ")
}
