package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AgencyTokenProject struct {

	// 委托方A的项目名称。
	Name string `json:"name" xml:"name"`

	// 委托方A的项目ID。
	Id string `json:"id" xml:"id"`

	Domain *AgencyTokenProjectDomain `json:"domain" xml:"domain"`
}

func (o AgencyTokenProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyTokenProject struct{}"
	}

	return strings.Join([]string{"AgencyTokenProject", string(data)}, " ")
}
