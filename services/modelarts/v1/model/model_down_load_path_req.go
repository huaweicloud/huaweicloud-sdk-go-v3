package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownLoadPathReq 下载临时路径请求
type DownLoadPathReq struct {

	// **参数解释**：带授权信息的URL的过期时间。 **约束限制**：0-86400。 **取值范围**：不涉及。 **默认取值**：3000。
	Expires *int64 `json:"expires,omitempty"`

	// **参数解释**：要下载文件的对象名。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Path *string `json:"path,omitempty"`
}

func (o DownLoadPathReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownLoadPathReq struct{}"
	}

	return strings.Join([]string{"DownLoadPathReq", string(data)}, " ")
}
