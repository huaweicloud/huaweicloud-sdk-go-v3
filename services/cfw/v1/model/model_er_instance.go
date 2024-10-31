package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ErInstance struct {

	// ER ID，创建ER时产生的ID
	Id *string `json:"id,omitempty"`

	// ER名称
	Name *string `json:"name,omitempty"`

	// 项目ID, 可以从调API处获取，也可以从控制台获取。[项目ID获取方式](cfw_02_0015.xml)
	ProjectId *string `json:"project_id,omitempty"`

	// 企业路由器连接id，该连接用于连接防火墙和企业路由器，此字段可在通过id在ER界面查询指定er后在管理连接界面查询连接了解连接具体情况。
	AttachmentId *string `json:"attachment_id,omitempty"`
}

func (o ErInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErInstance struct{}"
	}

	return strings.Join([]string{"ErInstance", string(data)}, " ")
}
