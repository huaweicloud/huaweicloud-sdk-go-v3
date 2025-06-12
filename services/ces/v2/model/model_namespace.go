package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Namespace 服务的命名空间，查询各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”
type Namespace struct {
}

func (o Namespace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Namespace struct{}"
	}

	return strings.Join([]string{"Namespace", string(data)}, " ")
}
