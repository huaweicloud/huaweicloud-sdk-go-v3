package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTempApiKeyReq 创建临时apike请求体
type CreateTempApiKeyReq struct {

	// **参数解释：** 过期时间。 **约束限制：** 不能是小数。 **取值范围：** 最少1小时，最多24小时。 **默认取值：** 不涉及。
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// **参数解释**：工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc)未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o CreateTempApiKeyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTempApiKeyReq struct{}"
	}

	return strings.Join([]string{"CreateTempApiKeyReq", string(data)}, " ")
}
