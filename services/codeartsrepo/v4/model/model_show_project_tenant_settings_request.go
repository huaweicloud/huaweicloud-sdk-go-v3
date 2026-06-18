package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectTenantSettingsRequest Request Object
type ShowProjectTenantSettingsRequest struct {

	// **参数解释：** 项目的32位uuid，项目唯一标识，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/en-us/api-projectman/ListProjectsV4.html)](tag:hws_hk)[[查询项目列表](https://support.huaweicloud.com/intl/zh-cn/api-projectman/ListProjectsV4.html)](tag:hws_hk_ch)[[查询项目列表](https://support.huaweicloud.com/eu/api-projectman/ListProjectsV4.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** project_id传参时表示查询的是项目所有者对应租户的设置 **取值范围：** 字符串长度32。
	ProjectId *string `json:"project_id,omitempty"`
}

func (o ShowProjectTenantSettingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectTenantSettingsRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectTenantSettingsRequest", string(data)}, " ")
}
