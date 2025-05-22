package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecycleBinResponseBody struct {

	// 参数解释：项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	Policy *RecycleBinPolicy `json:"policy,omitempty"`

	// 是否启用回收站。  取值： - true：启用回收站。 - false：不启用回收站。
	Enable *bool `json:"enable,omitempty"`
}

func (o RecycleBinResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleBinResponseBody struct{}"
	}

	return strings.Join([]string{"RecycleBinResponseBody", string(data)}, " ")
}
