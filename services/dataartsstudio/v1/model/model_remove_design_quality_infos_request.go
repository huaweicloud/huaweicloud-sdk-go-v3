package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveDesignQualityInfosRequest Request Object
type RemoveDesignQualityInfosRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 表的ID，填写String类型替代Long类型。
	TableId string `json:"table_id"`

	// 表类型，默认值是业务表。TABLE_MODEL(业务表(逻辑实体/物理表))、AGGREGATION_LOGIC_TABLE(汇总表)、FACT_LOGIC_TABLE(事实表)、DIMENSION_LOGIC_TABLE(维度表)。 - TABLE_MODEL - AGGREGATION_LOGIC_TABLE - FACT_LOGIC_TABLE - DIMENSION_LOGIC_TABLE
	TableType string `json:"table_type"`
}

func (o RemoveDesignQualityInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveDesignQualityInfosRequest struct{}"
	}

	return strings.Join([]string{"RemoveDesignQualityInfosRequest", string(data)}, " ")
}
