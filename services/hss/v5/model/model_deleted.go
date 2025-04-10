package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Deleted 是否删除，包含如下:   - true ：已删除   - false : 未删除
type Deleted struct {
}

func (o Deleted) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Deleted struct{}"
	}

	return strings.Join([]string{"Deleted", string(data)}, " ")
}
