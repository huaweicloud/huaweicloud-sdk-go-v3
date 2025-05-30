package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePolicyVersionReqBody the request body of create policy version request.
type CreatePolicyVersionReqBody struct {

	// 自定义身份策略或系统预置身份策略的策略文档的json格式。下面的字符`= < > ( ) |`是语法中的特殊字符，不包含在身份策略中。  问号`?`表示元素是可选的。例如`sid_block?`。  竖线`|`表示可选项，括号定义了可选项的范围。例如`(\"Allow\" | \"Deny\")`。  当一个元素允许多个值时，使用重复值、`,`以及`...`表示。例如`[ <policy_statement>, <policy_statement>, ... ]`。  下面的递归文法描述了身份策略的语法： ``` policy = {   <version_block>,   <statement_block> }  <version_block> = \"Version\" : (\"5.0\")  <statement_block> = \"Statement\" : [ <policy_statement>, <policy_statement>, ... ]  <policy_statement> = {   <sid_block?>,   <effect_block>,   <action_block>,   <resource_block?>,   <condition_block?> }  <sid_block> = \"Sid\" : <sid_string>  <effect_block> = \"Effect\" : (\"Allow\" | \"Deny\")  <action_block> = (\"Action\" | \"NotAction\") : [ <action_string>, <action_string>, ... ]  <resource_block> = (\"Resource\" | \"NotResource\") : [ <resource_string>, <resource_string>, ... ]  <condition_block> = \"Condition\" : { <condition_map> }  <condition_map> = {   <condition_type_string> : { <condition_key_string> : <condition_value_list> },   <condition_type_string> : { <condition_key_string> : <condition_value_list> },   ... }  <condition_value_list> = ( <condition_value> | [ <condition_value>, <condition_value>, ... ] )  <condition_value> = \"string\" ```
	PolicyDocument string `json:"policy_document"`

	// 是否设置为默认版本。
	SetAsDefault *bool `json:"set_as_default,omitempty"`
}

func (o CreatePolicyVersionReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyVersionReqBody struct{}"
	}

	return strings.Join([]string{"CreatePolicyVersionReqBody", string(data)}, " ")
}
