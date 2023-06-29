package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVmsSendTaskRequestBody 生成发送短信请求体。
type CreateVmsSendTaskRequestBody struct {

	// 任务名称。
	TaskName string `json:"task_name"`

	// 智能信息基础版模板ID。
	TplId string `json:"tpl_id"`

	// 失效时间（小时，范围是1~72小时）。
	ExpirationTime *int32 `json:"expiration_time,omitempty"`

	// 群发手机号码列表，最多支持5000个号码。  > 长度指的是单个号码的长度。 > mobiles和individual_params字段只能二选一。
	Mobiles []string `json:"mobiles"`

	// 群发动态参数数组。 - 参数顺序按照模板创建时参数占位符的顺序传入，例如创建模板时设置动参有#p_1#、#p_2#、#p_3#，则传入的参数数组顺序第一个元素为#p_1#，第二个元素是#p_2#，第三个元素为#p_3#。 - mobiles不填时，此字段被忽略。
	DyncParams *[]ContentParam `json:"dync_params,omitempty"`

	// 个性化手机号码及动态参数数组。  mobiles和individual_params字段只能二选一。
	IndividualParams *[]IndividualParam `json:"individual_params,omitempty"`

	// 智能信息基础版扩展字段。
	Exdata *string `json:"exdata,omitempty"`
}

func (o CreateVmsSendTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVmsSendTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVmsSendTaskRequestBody", string(data)}, " ")
}
