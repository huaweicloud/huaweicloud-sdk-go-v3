package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadModelServiceTaskFileResponse Response Object
type UploadModelServiceTaskFileResponse struct {

	// **参数解释**： 上传文件路径。 **约束限制**： 不涉及 **取值范围**： 长度为[0-1024]个字符。 **默认取值**： 不涉及
	Path           *string `json:"path,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadModelServiceTaskFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadModelServiceTaskFileResponse struct{}"
	}

	return strings.Join([]string{"UploadModelServiceTaskFileResponse", string(data)}, " ")
}
