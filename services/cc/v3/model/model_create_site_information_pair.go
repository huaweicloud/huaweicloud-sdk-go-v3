package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSiteInformationPair struct {

	// 端到端(P2P)类型分支网络连接的两个端点定义，长度固定为2的数组。
	Sites []CreateSiteInformation `json:"sites"`
}

func (o CreateSiteInformationPair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSiteInformationPair struct{}"
	}

	return strings.Join([]string{"CreateSiteInformationPair", string(data)}, " ")
}
