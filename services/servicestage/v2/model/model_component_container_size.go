package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentContainerSize struct {

	// 资源规格，可查询获取所有支持的应用资源规格接口获取系统预定义好的资源规格。
	Id *string `json:"id,omitempty"`
}

func (o ComponentContainerSize) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentContainerSize struct{}"
	}

	return strings.Join([]string{"ComponentContainerSize", string(data)}, " ")
}
