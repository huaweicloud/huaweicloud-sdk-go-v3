package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageCreateReq struct {

	// **参数解释**：该镜像所对应的描述信息。 **约束限制**：不涉及。 **取值范围**：长度限制512个字符。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：镜像名称。 **约束限制**：不涉及。 **取值范围**：长度限制为128个字符，支持小写字母、数字、中划线、下划线和点，字符串必须以小写字母或数字开头和结尾。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：镜像所属组织，可以在SWR控制台“组织管理”创建和查看。 **约束限制**：不涉及。 **取值范围**：长度限制为64个字符，支持大小写字母、数字、中划线、下划线和点号，且必须是小写字母开头。 **默认取值**：不涉及。
	Namespace string `json:"namespace"`

	// **参数解释**：镜像tag。 **约束限制**：不涉及。 **取值范围**：长度限制64个字符，支持大小写字母、数字、中划线、下划线和点号。 **默认取值**：不涉及。
	Tag *string `json:"tag,omitempty"`

	// **参数解释**：工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc)未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：0。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：企业版SWR仓库ID。 **参数约束**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SwrInstanceId *string `json:"swr_instance_id,omitempty"`

	// **参数解释**：企业版SWR仓库域名。 **参数约束**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SwrInstanceDomain *string `json:"swr_instance_domain,omitempty"`
}

func (o ImageCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageCreateReq struct{}"
	}

	return strings.Join([]string{"ImageCreateReq", string(data)}, " ")
}
