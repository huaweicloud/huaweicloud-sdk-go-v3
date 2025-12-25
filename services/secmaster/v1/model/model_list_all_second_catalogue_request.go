package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllSecondCatalogueRequest Request Object
type ListAllSecondCatalogueRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 目录类型（内置或自定义）
	CatalogueType *string `json:"catalogue_type,omitempty"`

	// 目录编码
	CatalogueCode *string `json:"catalogue_code,omitempty"`
}

func (o ListAllSecondCatalogueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllSecondCatalogueRequest struct{}"
	}

	return strings.Join([]string{"ListAllSecondCatalogueRequest", string(data)}, " ")
}
