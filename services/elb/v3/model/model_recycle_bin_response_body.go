package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecycleBinResponseBody **参数解释**：回收站响应体。
type RecycleBinResponseBody struct {

	// **参数解释**：项目ID。获取方式请参见[获取项目ID](elb_fl_0008.xml)。  **取值范围**：长度为32个字符，由小写字母和数字组成。
	ProjectId *string `json:"project_id,omitempty"`

	Policy *RecycleBinPolicy `json:"policy,omitempty"`

	// **参数解释**：是否启用回收站。  **取值范围**： - true：启用回收站。 - false：不启用回收站。
	Enable *bool `json:"enable,omitempty"`
}

func (o RecycleBinResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleBinResponseBody struct{}"
	}

	return strings.Join([]string{"RecycleBinResponseBody", string(data)}, " ")
}
