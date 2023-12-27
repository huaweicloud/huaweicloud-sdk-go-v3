package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdePrivilageProjectInfo struct {

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// tab_id集合
	Ids *[]string `json:"ids,omitempty"`
}

func (o IdePrivilageProjectInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdePrivilageProjectInfo struct{}"
	}

	return strings.Join([]string{"IdePrivilageProjectInfo", string(data)}, " ")
}
