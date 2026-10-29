package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OldTaskIdInResponse **参数解释：** 旧版任务ID。即新建检查任务接口的返回值，每个代码检查任务及其下的分支任务、通过API创建的增量任务都有唯一的任务ID。 **取值范围：** 不涉及。
type OldTaskIdInResponse struct {
}

func (o OldTaskIdInResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OldTaskIdInResponse struct{}"
	}

	return strings.Join([]string{"OldTaskIdInResponse", string(data)}, " ")
}
