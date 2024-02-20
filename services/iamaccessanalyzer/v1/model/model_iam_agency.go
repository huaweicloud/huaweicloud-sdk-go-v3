package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IamAgency struct {

	// 此策略的json格式策略文档。
	TrustPolicy string `json:"trust_policy"`
}

func (o IamAgency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IamAgency struct{}"
	}

	return strings.Join([]string{"IamAgency", string(data)}, " ")
}
