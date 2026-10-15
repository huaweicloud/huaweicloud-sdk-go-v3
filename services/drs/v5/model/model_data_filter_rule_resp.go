package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataFilterRuleResp 数据过滤规则响应体
type DataFilterRuleResp struct {

	// 数据过滤规则批次ID，标识同一批次内的多个过滤对象属于同一个规则配置。
	BatchId *int64 `json:"batch_id,omitempty"`

	// 过滤对象信息列表，列出此规则下所有需进行数据过滤的具体数据库对象
	FilterObjectList *[]DataFilterObjectResp `json:"filter_object_list,omitempty"`

	FilterRulesResp *GeneralDataFilterResp `json:"filter_rules_resp,omitempty"`

	AdvancedSettingResp *DataFilteringAdvancedSetting `json:"advanced_setting_resp,omitempty"`
}

func (o DataFilterRuleResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataFilterRuleResp struct{}"
	}

	return strings.Join([]string{"DataFilterRuleResp", string(data)}, " ")
}
