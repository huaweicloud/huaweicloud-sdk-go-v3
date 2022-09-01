package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// project信息
type ProjectInfo struct {
	Domain *DomainInfo `json:"domain,omitempty" xml:"domain"`

	// project id
	Id *string `json:"id,omitempty" xml:"id"`

	// project name
	Name string `json:"name" xml:"name"`
}

func (o ProjectInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectInfo struct{}"
	}

	return strings.Join([]string{"ProjectInfo", string(data)}, " ")
}
