package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DomainProjectsInfo
type DomainProjectsInfo struct {

	//
	ProjectId string `json:"project_id"`

	//
	ProjectName string `json:"project_name"`
}

func (o DomainProjectsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainProjectsInfo struct{}"
	}

	return strings.Join([]string{"DomainProjectsInfo", string(data)}, " ")
}
