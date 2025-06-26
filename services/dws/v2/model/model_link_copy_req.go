package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LinkCopyReq **参数解释**： 快照复制请求。 **取值范围**： 不涉及。
type LinkCopyReq struct {

	// **参数解释**： 快照名称。 **取值范围**： 不涉及。
	BackupName string `json:"backup_name"`

	// **参数解释**： 描述信息。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o LinkCopyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinkCopyReq struct{}"
	}

	return strings.Join([]string{"LinkCopyReq", string(data)}, " ")
}
