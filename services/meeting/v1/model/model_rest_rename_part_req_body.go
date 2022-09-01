package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 重命名会场消息体。
type RestRenamePartReqBody struct {

	// 与会者标识。 已入会的必须填写该字段。
	ParticipantID *string `json:"participantID,omitempty" xml:"participantID"`

	// 与会者号码。
	Number string `json:"number" xml:"number"`

	// 新名字。
	NewName string `json:"newName" xml:"newName"`
}

func (o RestRenamePartReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestRenamePartReqBody struct{}"
	}

	return strings.Join([]string{"RestRenamePartReqBody", string(data)}, " ")
}
