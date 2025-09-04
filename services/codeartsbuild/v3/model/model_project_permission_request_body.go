package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectPermissionRequestBody 创建更新分组返回体
type ProjectPermissionRequestBody struct {

	// **参数解释**： 填写需要查询构建历史列表的构建任务ID。获取方法：在构建任务详情页，拷贝浏览器URL末尾的32位数字、字母组合的字符串，即为构建任务ID。 **约束限制**： 不涉及。 **取值范围**： 只能是英文字母和数字，长度为32个字符。 **默认取值**： 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**： CodeArts项目ID。获取方式请参考[获取CodeArts项目ID](https://support.huaweicloud.com/api-codeci/cloudbuild_03_0022.html)。 **约束限制**： 不涉及。 **取值范围**： 32位数字、字母组合的字符串。 **默认取值**： 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 是否同步最新一次项目权限。 **约束限制**： 不涉及。 **取值范围**： ● true：使用项目级权限。 ● false：不使用项目级权限。 **默认取值**： 不涉及。
	ProjectSwitch *bool `json:"project_switch,omitempty"`
}

func (o ProjectPermissionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectPermissionRequestBody struct{}"
	}

	return strings.Join([]string{"ProjectPermissionRequestBody", string(data)}, " ")
}
