package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowFileTreeResult struct {

	// **参数解释**: 上传权限。 **取值范围**: true：有权限。 false：无权限。
	UploadAccess *string `json:"uploadAccess,omitempty"`

	// **参数解释**: 总数。 **取值范围**: 不涉及。
	Total *string `json:"total,omitempty"`

	// **参数解释**: 父路径。 **取值范围**: 不涉及。
	Uri *string `json:"uri,omitempty"`

	// **参数解释**: 创建时间。 **取值范围**: 不涉及。
	Created *string `json:"created,omitempty"`

	// **参数解释**: 子文件及文件夹信息。 **取值范围**: 不涉及。
	Children *[]ShowFileTreeResultChildren `json:"children,omitempty"`
}

func (o ShowFileTreeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFileTreeResult struct{}"
	}

	return strings.Join([]string{"ShowFileTreeResult", string(data)}, " ")
}
