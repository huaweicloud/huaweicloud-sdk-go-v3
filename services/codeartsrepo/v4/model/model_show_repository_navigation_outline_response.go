package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepositoryNavigationOutlineResponse Response Object
type ShowRepositoryNavigationOutlineResponse struct {

	// **参数解释：** 结果标识。 **约束限制：** 不涉及。
	Result *string `json:"result,omitempty"`

	// **参数解释：** 结果消息。 **约束限制：** 不涉及。
	Message *string `json:"message,omitempty"`

	// **参数解释：** 文件路径。 **约束限制：** 不涉及。
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释：** 所在版本号（commit id）。 **约束限制：** 不涉及。
	Revision *string `json:"revision,omitempty"`

	// **参数解释：** 符号列表。 **约束限制：** 不涉及。
	Symbols        *[]SymbolNodeDto `json:"symbols,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowRepositoryNavigationOutlineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryNavigationOutlineResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryNavigationOutlineResponse", string(data)}, " ")
}
