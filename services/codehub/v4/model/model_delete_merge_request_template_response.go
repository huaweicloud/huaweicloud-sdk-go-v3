package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMergeRequestTemplateResponse Response Object
type DeleteMergeRequestTemplateResponse struct {

	// **参数解释：** 合并请求模板主键id
	Id             *int32 `json:"id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteMergeRequestTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMergeRequestTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteMergeRequestTemplateResponse", string(data)}, " ")
}
