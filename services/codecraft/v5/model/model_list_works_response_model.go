package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListWorksResponseModel struct {

	// 作品ID
	WorksId string `json:"works_id"`

	// 租户ID
	DomainId string `json:"domain_id"`

	// 团队ID
	TeamId string `json:"team_id"`
}

func (o ListWorksResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorksResponseModel struct{}"
	}

	return strings.Join([]string{"ListWorksResponseModel", string(data)}, " ")
}
