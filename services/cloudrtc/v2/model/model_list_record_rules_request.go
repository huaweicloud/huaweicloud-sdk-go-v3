package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecordRulesRequest Request Object
type ListRecordRulesRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 应用id
	AppId string `json:"app_id"`

	// 查询结果起始编号，此处代表分页的页码
	Offset *int32 `json:"offset,omitempty"`

	// 查询结果集数量，此处代表每一页的条数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRecordRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordRulesRequest struct{}"
	}

	return strings.Join([]string{"ListRecordRulesRequest", string(data)}, " ")
}
