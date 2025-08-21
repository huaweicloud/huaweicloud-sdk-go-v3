package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGroupRequest Request Object
type DeleteGroupRequest struct {

	// **参数解释：** 项目的32位uuid，项目唯一标识，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/en-us/api-projectman/ListProjectsV4.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **取值范围：** 字符串长度32。
	ProjectId string `json:"project_id"`

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`
}

func (o DeleteGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteGroupRequest", string(data)}, " ")
}
