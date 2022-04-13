package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 操作配置信息。
type UpdateYmlsReqEdit struct {
	Modify *UpdateYmlsReqEditModify `json:"modify"`
}

func (o UpdateYmlsReqEdit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateYmlsReqEdit struct{}"
	}

	return strings.Join([]string{"UpdateYmlsReqEdit", string(data)}, " ")
}
