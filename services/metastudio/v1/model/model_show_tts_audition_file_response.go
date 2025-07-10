package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTtsAuditionFileResponse Response Object
type ShowTtsAuditionFileResponse struct {

	// 试听文件是否已生成完成。该标记为false时，应该每隔5秒再次调用本接口获取试听文件。当存在该参数时，将会返回以下message和files两个字段信息
	IsFileComplete *bool `json:"is_file_complete,omitempty"`

	// 异常信息。
	Message *string `json:"message,omitempty"`

	// 试听文件列表。
	Files          *[]AuditionFile `json:"files,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowTtsAuditionFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTtsAuditionFileResponse struct{}"
	}

	return strings.Join([]string{"ShowTtsAuditionFileResponse", string(data)}, " ")
}
