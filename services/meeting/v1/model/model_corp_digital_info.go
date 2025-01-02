package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CorpDigitalInfo 数字资产信息
type CorpDigitalInfo struct {

	// 音色资产标识ID。
	Account *string `json:"account,omitempty"`

	// 当前会议企业ID。
	CorpId *string `json:"corpId,omitempty"`

	// 类型，PUBLIC：公共、LOCAL：本地会议、PRIVATE：克隆专用。
	Type *string `json:"type,omitempty"`

	// 音色名称。
	Name *string `json:"name,omitempty"`
}

func (o CorpDigitalInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CorpDigitalInfo struct{}"
	}

	return strings.Join([]string{"CorpDigitalInfo", string(data)}, " ")
}
