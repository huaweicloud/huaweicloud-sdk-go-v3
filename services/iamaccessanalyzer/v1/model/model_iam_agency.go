package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IamAgency IAM信任委托。
type IamAgency struct {

	// 该策略JSON格式策略文档。
	TrustPolicy string `json:"trust_policy"`
}

func (o IamAgency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IamAgency struct{}"
	}

	return strings.Join([]string{"IamAgency", string(data)}, " ")
}
