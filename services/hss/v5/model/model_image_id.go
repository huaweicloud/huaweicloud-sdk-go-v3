package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageId **参数解释**： 镜像ID **取值范围**： 字符长度1-64位
type ImageId struct {
}

func (o ImageId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageId struct{}"
	}

	return strings.Join([]string{"ImageId", string(data)}, " ")
}
