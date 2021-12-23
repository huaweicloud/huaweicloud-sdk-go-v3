package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// extend对象
type Extend struct {
	// 扩容后文件系统的新容量，以GB为单位。扩容步长大于等于100GB。  普通文件系统容量，取值范围500~32768。  带宽型文件系统，容量范围是10240~327680

	NewSize int32 `json:"new_size"`
}

func (o Extend) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Extend struct{}"
	}

	return strings.Join([]string{"Extend", string(data)}, " ")
}
