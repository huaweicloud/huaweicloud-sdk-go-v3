package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 重命名与会者请求。
type RestRenamePartReqBody struct {

	// 与会者标识。 已入会的必须填写该字段。
	ParticipantID *string `json:"participantID,omitempty"`

	// 与会者号码。
	Number string `json:"number"`

	// 新名称。
	NewName string `json:"newName"`
}

func (o RestRenamePartReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestRenamePartReqBody struct{}"
	}

	return strings.Join([]string{"RestRenamePartReqBody", string(data)}, " ")
}
