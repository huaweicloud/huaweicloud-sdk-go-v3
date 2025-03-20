package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyEx 委托或信任委托。
type AgencyEx struct {

	// 统一资源名称。
	Urn string `json:"urn"`

	// 信任委托信任策略的策略文档的json格式。下面的字符`= < > ( ) |`是语法中的特殊字符，不包含在信任策略中。  问号`?`表示元素是可选的。例如`sid_block?`。  竖线`|`表示可选项，括号定义了可选项的范围。例如`(\"Allow\" | \"Deny\")`。  当一个元素允许多个值时，使用重复值、`,`以及`...`表示。例如`[ <policy_statement>, <policy_statement>, ... ]`。  下面的递归文法描述了信任策略的语法： ``` policy = {   <version_block>,   <statement_block> }  <version_block> = \"Version\" : (\"5.0\")  <statement_block> = \"Statement\" : [ <policy_statement>, <policy_statement>, ... ]  <policy_statement> = {   <sid_block?>,   <principal_block>,   <effect_block>,   <action_block>,   <resource_block?>,   <condition_block?> }  <sid_block> = \"Sid\" : <sid_string>  <principal_block> = (\"Principal\" | \"NotPrincipal\") : <principal_map>  <principal_map> = { <principal_map_entry>, <principal_map_entry>, ... }  <principal_map_entry> = (\"IAM\" | \"Service\") : [ <principal_id_string>, <principal_id_string>, ... ]  <effect_block> = \"Effect\" : (\"Allow\" | \"Deny\")  <action_block> = (\"Action\" | \"NotAction\") : [ <action_string>, <action_string>, ... ]  <resource_block> = (\"Resource\" | \"NotResource\") : [ <resource_string>, <resource_string>, ... ]  <condition_block> = \"Condition\" : { <condition_map> }  <condition_map> = {   <condition_type_string> : { <condition_key_string> : <condition_value_list> },   <condition_type_string> : { <condition_key_string> : <condition_value_list> },   ... }  <condition_value_list> = ( <condition_value> | [ <condition_value>, <condition_value>, ... ] )  <condition_value> = \"string\" ```
	TrustPolicy *string `json:"trust_policy,omitempty"`

	// 委托或信任委托创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 委托或信任委托描述信息。
	Description *string `json:"description,omitempty"`

	// 委托或信任委托最大会话时长，默认为3600秒。
	MaxSessionDuration int32 `json:"max_session_duration"`

	// 资源路径，默认为空串。由若干段字符串拼接而成，每段先包含一个或多个字母、数字、\".\"、\",\"、\"+\"、\"@\"、\"=\"、\"_\"或\"-\"，并以\"/\"结尾，例如\"foo/bar/\"。
	Path string `json:"path"`

	// 委托或信任委托ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	AgencyId string `json:"agency_id"`

	// 委托或信任委托名称，长度为1到64个字符，只包含字母、数字、\"_\"、\"+\"、\"=\"、\",\"、\".\"、\"@\"和\"-\"的字符串。
	AgencyName string `json:"agency_name"`

	// 被委托方账号ID，仅存在于委托中，不存在于信任委托中。
	TrustDomainId *string `json:"trust_domain_id,omitempty"`

	// 被委托方账号名，仅存在于委托中，不存在于信任委托中。
	TrustDomainName *string `json:"trust_domain_name,omitempty"`

	// 自定义标签列表。
	Tags []Tag `json:"tags"`
}

func (o AgencyEx) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyEx struct{}"
	}

	return strings.Join([]string{"AgencyEx", string(data)}, " ")
}
