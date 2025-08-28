package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPluginInstallScriptResponse Response Object
type ListPluginInstallScriptResponse struct {

	// **参数解释**: docker插件脚本信息 **取值范围**: 最小值0，最大值200
	DataList       *[]PluginInstallScriptResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListPluginInstallScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginInstallScriptResponse struct{}"
	}

	return strings.Join([]string{"ListPluginInstallScriptResponse", string(data)}, " ")
}
