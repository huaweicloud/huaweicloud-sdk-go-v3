package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Er struct {

	// ER ID，创建东西向防护引用的 ID
	ErId *string `json:"er_id,omitempty"`

	// 企业路由器连接id，该连接用于连接防火墙和企业路由器，此字段可在通过id在ER界面查询指定er后在管理连接界面查询连接了解连接具体情况。
	ErAttachId *string `json:"er_attach_id,omitempty"`
}

func (o Er) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Er struct{}"
	}

	return strings.Join([]string{"Er", string(data)}, " ")
}
