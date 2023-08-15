package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectDto 项目
type ProjectDto struct {

	// 项目id
	Id *string `json:"id,omitempty"`

	// 项目名
	Name *string `json:"name,omitempty"`
}

func (o ProjectDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectDto struct{}"
	}

	return strings.Join([]string{"ProjectDto", string(data)}, " ")
}
