package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SpecStrideNum 资源步数，最小为10000，步长为1。输入范围还需要满足规格列表接口的步长白名单
type SpecStrideNum struct {
}

func (o SpecStrideNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecStrideNum struct{}"
	}

	return strings.Join([]string{"SpecStrideNum", string(data)}, " ")
}
