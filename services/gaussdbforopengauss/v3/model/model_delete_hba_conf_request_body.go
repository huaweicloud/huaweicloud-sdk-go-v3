package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteHbaConfRequestBody struct {

	// **参数解释**: 需要删除的hba配置信息。 **取值范围**: 不涉及。
	HbaConfs []HbaConfOption `json:"hba_confs"`
}

func (o DeleteHbaConfRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHbaConfRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteHbaConfRequestBody", string(data)}, " ")
}
