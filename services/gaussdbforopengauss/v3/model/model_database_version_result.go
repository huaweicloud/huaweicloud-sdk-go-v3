package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatabaseVersionResult struct {

	// **参数解释**： 数据库三位引擎版本。 **取值范围**： 不涉及。
	SoftwareVersion string `json:"software_version"`

	// **参数解释**： 数据库三位引擎版本对应的热补丁信息。 **取值范围**： 不涉及。
	Hotfixes []DbHotfixInfoResult `json:"hotfixes"`
}

func (o DatabaseVersionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseVersionResult struct{}"
	}

	return strings.Join([]string{"DatabaseVersionResult", string(data)}, " ")
}
