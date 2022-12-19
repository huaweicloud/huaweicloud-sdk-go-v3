package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateEdgeAppResponse struct {

	// **参数说明**：用户自定义应用唯一ID。  **取值范围**：只允许字母、数字、下划线（_）、连接符（-）、美元符号（$）的组合。
	EdgeAppId *string `json:"edge_app_id,omitempty"`

	// **参数说明**：应用描述。  **取值范围**：只允许中文、字母、数字、下划线（_）、中文分号（；）、中文冒号（：）、中文问号（？）、中文感叹号（！）中文逗号（，）、中文句号（。）、英文引号（;）、英文冒号（:）、英文逗号（,）、英文句号（.）、英文问号（?）、英文感叹号（!）、顿号（、）、连接符（-）的组合。
	Description *string `json:"description,omitempty"`

	// **参数说明**：创建时间。
	CreatedTime *string `json:"created_time,omitempty"`

	// **参数说明**：更新时间。
	LastModifiedTime *string `json:"last_modified_time,omitempty"`

	// **参数说明**：最新发布版本。
	LastPublishedVersion *string `json:"last_published_version,omitempty"`
	HttpStatusCode       int     `json:"-"`
}

func (o UpdateEdgeAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeAppResponse struct{}"
	}

	return strings.Join([]string{"UpdateEdgeAppResponse", string(data)}, " ")
}
