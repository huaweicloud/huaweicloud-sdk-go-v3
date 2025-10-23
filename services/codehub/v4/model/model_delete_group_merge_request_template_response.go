package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGroupMergeRequestTemplateResponse Response Object
type DeleteGroupMergeRequestTemplateResponse struct {

	// **参数解释：** 合并请求模板主键id
	Id             *int32 `json:"id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteGroupMergeRequestTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGroupMergeRequestTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteGroupMergeRequestTemplateResponse", string(data)}, " ")
}
