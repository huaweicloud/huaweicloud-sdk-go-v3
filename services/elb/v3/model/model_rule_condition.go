package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type RuleCondition struct {

	// 匹配项的名称。[该字段固定为空字符串](tag:dt,dt_test)   [当type为HOST_NAME、PATH、METHOD、SOURCE_IP时，该字段固定为空字符串。   当type为HEADER时，key表示请求头参数的名称，value表示请求头参数的值。key的长度限制1-40字符，只允许包含字母、数字和-_。   当type为QUERY_STRING时，key表示查询参数的名称，value表示查询参数的值。key的长度限制为1-128字符，不支持空格，中括号，大括号，尖括号，反斜杠，双引号，'#'，'&'，'|'，‘%’，‘~’，字母区分大小写。](tag:hws,hws_hk,ocb,tlf,ctc,hcs,sbc,g42,tm,cmcc,hk_g42,mix,hk_sbc,hws_ocb,fcs)
	Key string `json:"key"`

	// 匹配项的值。   当type为HOST_NAME时，key固定为空字符串，value表示域名的值。value长度1-128字符，字符串只能包含英文字母、数字、“-”、“.”或“*”，必须以字母、数字或“*”开头，“*”只能出现在开头且必须以*.开始。   当type为PATH时，key固定为空字符串，value表示请求路径的值。value长度1-128字符。当转发规则的compare_type为STARTS_WITH、EQUAL_TO时，字符串只能包含英文字母、数字、_~';@^-%#&$.*+?,=!:|\\/()\\[\\]{}，且必须以\"/\"开头。   [当type为HEADER时，key表示请求头参数的名称，value表示请求头参数的值。value长度限制1-128字符，不支持空格，双引号，支持以下通配符：*（匹配0个或更多字符）和？（正好匹配1个字符）。   当type为QUERY_STRING时，key表示查询参数的名称，value表示查询参数的值。value长度限制为1-128字符，不支持空格，中括号，大括号，尖括号，反斜杠，双引号，'#'，'&'，'|'，‘%’，‘~’，字母区分大小写，支持通配符：*（匹配0个或更多字符）和？（正好匹配1个字符）   当type为METHOD时，key固定为空字符串，value表示请求方式。value取值范围为：GET, PUT, POST, DELETE, PATCH, HEAD, OPTIONS。   当type为SOURCE_IP时，key固定为空字符串，value表示请求源地址。value为CIDR格式，支持ipv4，ipv6。 例如192.168.0.2/32，2049::49/64。](tag:hws,hws_hk,ocb,tlf,ctc,hcs,sbc,g42,tm,cmcc,hk_g42,mix,hk_sbc,hws_ocb,fcs)
	Value string `json:"value"`
}

func (o RuleCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleCondition struct{}"
	}

	return strings.Join([]string{"RuleCondition", string(data)}, " ")
}
