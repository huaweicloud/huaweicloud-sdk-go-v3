package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExecuteUploadPptResponse struct {

	// 图片链接列表
	List           *[]ImageUrlResp `json:"list,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ExecuteUploadPptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteUploadPptResponse struct{}"
	}

	return strings.Join([]string{"ExecuteUploadPptResponse", string(data)}, " ")
}
