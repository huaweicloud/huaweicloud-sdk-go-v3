package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NfsResp nfs方式的挂载卷。
type NfsResp struct {

	// **参数解释**：nfs服务端路径，如：“10.10.10.10:/example/path”。 **取值范围**：不涉及。
	NfsServerPath *string `json:"nfs_server_path,omitempty"`

	// **参数解释**：挂载到训练容器中的路径，如：“/example/path”。 **取值范围**：不涉及。
	LocalPath *string `json:"local_path,omitempty"`

	// **参数解释**：nfs挂载卷在容器中是否只读。 **取值范围**： - true：只读 - false：非只读
	ReadOnly *bool `json:"read_only,omitempty"`
}

func (o NfsResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NfsResp struct{}"
	}

	return strings.Join([]string{"NfsResp", string(data)}, " ")
}
