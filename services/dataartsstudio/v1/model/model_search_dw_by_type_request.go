package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchDwByTypeRequest Request Object
type SearchDwByTypeRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 是否查询最新的。
	ForceRefresh *bool `json:"force_refresh,omitempty"`

	// 数据连接类型。
	DwType string `json:"dw_type"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// limit
	Offset *int32 `json:"offset,omitempty"`
}

func (o SearchDwByTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchDwByTypeRequest struct{}"
	}

	return strings.Join([]string{"SearchDwByTypeRequest", string(data)}, " ")
}
