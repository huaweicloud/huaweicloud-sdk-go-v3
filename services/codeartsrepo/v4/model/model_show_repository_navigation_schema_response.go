package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepositoryNavigationSchemaResponse Response Object
type ShowRepositoryNavigationSchemaResponse struct {

	// **参数解释：** 结果标识。 **约束限制：** 不涉及。
	Result *string `json:"result,omitempty"`

	// **参数解释：** 结果消息。 **约束限制：** 不涉及。
	Message *string `json:"message,omitempty"`

	Schema         *SchemaDto `json:"schema,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowRepositoryNavigationSchemaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryNavigationSchemaResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryNavigationSchemaResponse", string(data)}, " ")
}
