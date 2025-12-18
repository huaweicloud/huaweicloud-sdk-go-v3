package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProjectNoteRequiredAttributesRequest Request Object
type UpdateProjectNoteRequiredAttributesRequest struct {

	// **参数解释：** 项目的32位uuid，项目唯一标识，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/en-us/api-projectman/ListProjectsV4.html)](tag:hws_hk)[[查询项目列表](https://support.huaweicloud.com/eu/api-projectman/ListProjectsV4.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **取值范围：** 字符串长度32。
	ProjectId string `json:"project_id"`

	Body *PostNoteRequiredAttributeDto `json:"body,omitempty"`
}

func (o UpdateProjectNoteRequiredAttributesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectNoteRequiredAttributesRequest struct{}"
	}

	return strings.Join([]string{"UpdateProjectNoteRequiredAttributesRequest", string(data)}, " ")
}
