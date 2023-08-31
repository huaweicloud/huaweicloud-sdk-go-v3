package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EcnWithErRequest struct {

	// 企业路由器ID
	ErId string `json:"er_id"`

	// 企业路由器区域项目ID
	RegionProjectId string `json:"region_project_id"`

	// 区域ID
	RegionId string `json:"region_id"`
}

func (o EcnWithErRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EcnWithErRequest struct{}"
	}

	return strings.Join([]string{"EcnWithErRequest", string(data)}, " ")
}
