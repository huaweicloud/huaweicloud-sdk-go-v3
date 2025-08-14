package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageName **参数解释**： 镜像名称，只有容器类型的告警有 **取值范围**： 字符长度1-256位
type ImageName struct {
}

func (o ImageName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageName struct{}"
	}

	return strings.Join([]string{"ImageName", string(data)}, " ")
}
