package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ErInstance struct {

	// **参数解释**： ER ID，创建ER时产生的ID **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： ER名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 项目ID，可以通过调用API获取，也可以从控制台获取。[项目ID获取方式](cfw_02_0015.xml) **取值范围**： 不涉及
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 企业路由器连接id，该连接用于连接防火墙和企业路由器，此字段可在通过id在ER界面查询指定er后在管理连接界面查询连接了解连接具体情况。 **取值范围**： 不涉及
	AttachmentId *string `json:"attachment_id,omitempty"`
}

func (o ErInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErInstance struct{}"
	}

	return strings.Join([]string{"ErInstance", string(data)}, " ")
}
