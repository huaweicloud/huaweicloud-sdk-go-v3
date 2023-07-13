package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyMappingArray 用户（组）与IAM委托的映射关系数组结构体
type AgencyMappingArray struct {

	// 用户（组）与委托之间的映射关系详细信息。
	AgencyMappings []AgencyMapping `json:"agency_mappings"`
}

func (o AgencyMappingArray) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyMappingArray struct{}"
	}

	return strings.Join([]string{"AgencyMappingArray", string(data)}, " ")
}
