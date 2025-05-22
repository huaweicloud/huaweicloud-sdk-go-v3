package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeTypeDatastores **参数解释**： 内核版本信息。 **取值范围**： 数字、小数点，格式一般如 9.1.0、9.1.1.100。
type NodeTypeDatastores struct {

	// **参数解释**： 内核版本号。 **取值范围**： 数字、小数点，格式一般如 9.1.0、9.1.1.100。
	Version string `json:"version"`

	Attachments *NodeTypeDatastoresAttachments `json:"attachments,omitempty"`

	// **参数解释**： 版本类型。 **取值范围**： - STABLE：稳定版 - PREVIEW：预览版
	Role *string `json:"role,omitempty"`
}

func (o NodeTypeDatastores) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeTypeDatastores struct{}"
	}

	return strings.Join([]string{"NodeTypeDatastores", string(data)}, " ")
}
