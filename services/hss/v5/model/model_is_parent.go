package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsParent **参数解释**： 是否是父进程 **取值范围**： - true：是父进程 - false：不是父进程
type IsParent struct {
}

func (o IsParent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsParent struct{}"
	}

	return strings.Join([]string{"IsParent", string(data)}, " ")
}
