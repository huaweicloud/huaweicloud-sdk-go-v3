package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApigGroupDto struct {

	// 分组编号
	GroupId *string `json:"group_id,omitempty"`

	// 分组名称
	GroupName *string `json:"group_name,omitempty"`
}

func (o ApigGroupDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApigGroupDto struct{}"
	}

	return strings.Join([]string{"ApigGroupDto", string(data)}, " ")
}
