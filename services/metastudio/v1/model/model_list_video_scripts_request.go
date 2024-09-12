package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListVideoScriptsRequest Request Object
type ListVideoScriptsRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 按名称模糊查询。
	Name *string `json:"name,omitempty"`

	// 剧本类型。默认查询VIDEO_DRAFT。 * VIDEO_DRAFT：视频草稿。 * SYSTEM_VIDEO_TEMPLET： 系统视频模板。
	ScriptCatalog *ListVideoScriptsRequestScriptCatalog `json:"script_catalog,omitempty"`

	// 横竖屏类型（内部参数，不对外开放）。默认值是LANDSCAPE。 * LANDSCAPE：横屏。 * VERTICAL：竖屏。
	ViewMode *ListVideoScriptsRequestViewMode `json:"view_mode,omitempty"`
}

func (o ListVideoScriptsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVideoScriptsRequest struct{}"
	}

	return strings.Join([]string{"ListVideoScriptsRequest", string(data)}, " ")
}

type ListVideoScriptsRequestScriptCatalog struct {
	value string
}

type ListVideoScriptsRequestScriptCatalogEnum struct {
	VIDEO_DRAFT          ListVideoScriptsRequestScriptCatalog
	SYSTEM_VIDEO_TEMPLET ListVideoScriptsRequestScriptCatalog
}

func GetListVideoScriptsRequestScriptCatalogEnum() ListVideoScriptsRequestScriptCatalogEnum {
	return ListVideoScriptsRequestScriptCatalogEnum{
		VIDEO_DRAFT: ListVideoScriptsRequestScriptCatalog{
			value: "VIDEO_DRAFT",
		},
		SYSTEM_VIDEO_TEMPLET: ListVideoScriptsRequestScriptCatalog{
			value: "SYSTEM_VIDEO_TEMPLET",
		},
	}
}

func (c ListVideoScriptsRequestScriptCatalog) Value() string {
	return c.value
}

func (c ListVideoScriptsRequestScriptCatalog) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVideoScriptsRequestScriptCatalog) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListVideoScriptsRequestViewMode struct {
	value string
}

type ListVideoScriptsRequestViewModeEnum struct {
	LANDSCAPE ListVideoScriptsRequestViewMode
	VERTICAL  ListVideoScriptsRequestViewMode
}

func GetListVideoScriptsRequestViewModeEnum() ListVideoScriptsRequestViewModeEnum {
	return ListVideoScriptsRequestViewModeEnum{
		LANDSCAPE: ListVideoScriptsRequestViewMode{
			value: "LANDSCAPE",
		},
		VERTICAL: ListVideoScriptsRequestViewMode{
			value: "VERTICAL",
		},
	}
}

func (c ListVideoScriptsRequestViewMode) Value() string {
	return c.value
}

func (c ListVideoScriptsRequestViewMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVideoScriptsRequestViewMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
