package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProjectMergeRequestTemplateResponse Response Object
type DeleteProjectMergeRequestTemplateResponse struct {

	// **参数解释：** 合并请求模板主键id
	Id             *int32 `json:"id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteProjectMergeRequestTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProjectMergeRequestTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteProjectMergeRequestTemplateResponse", string(data)}, " ")
}
