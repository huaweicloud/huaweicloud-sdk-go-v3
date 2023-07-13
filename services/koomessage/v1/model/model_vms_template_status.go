package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VmsTemplateStatus 查询智能信息基础版模板返回体。
type VmsTemplateStatus struct {

	// 智能信息基础版模板ID。
	TplId *string `json:"tpl_id,omitempty"`

	// 智能信息基础版审核状态。 - 0：正常可用 - 1：审核中 - 2：审核不通过 - 3：模板已禁用 - 4：模板不存在 - 5：模板已过期
	AuditState *int32 `json:"audit_state,omitempty"`

	// 智能信息基础版模板状态的描述，若状态是审核不通过或被禁用，描述表示的是不通过或禁用的原因。 > 长度不超过 1024 字。
	AuditDesc *string `json:"audit_desc,omitempty"`

	// 智能信息基础版模板的大小。  >  单位：字节。
	TplSize *int32 `json:"tpl_size,omitempty"`

	// 模板截止有效日期，格式：yyyy-MM-ddTHH:mm:ssZ，0：表示永久有效。样例：2020-01-31T23:59:59Z。
	ValidTime *string `json:"valid_time,omitempty"`

	// 运营商的模板状态详情。
	StatusDetail *[]StatusDetail `json:"status_detail,omitempty"`

	// 智能信息基础版模板预览地址。
	PreviewUrl *string `json:"preview_url,omitempty"`

	// 智能信息基础版模板名称。
	TplName *string `json:"tpl_name,omitempty"`

	// 智能信息基础版模板主题。
	Title *string `json:"title,omitempty"`

	// 智能信息基础版模板签名。
	TplSign *string `json:"tpl_sign,omitempty"`

	// 智能信息基础版模板创建时间，格式：yyyy-MM-ddTHH:mm:ssZ。
	CreateTime *string `json:"create_time,omitempty"`

	// 智能信息基础版模板动参信息，静态模板没有动参，该字段填空。
	VarInfo *[]VarInfo `json:"var_info,omitempty"`

	// 资源分配标签，取值如下： - 三网一般 - 三网抗诉 - 单网一般 - 单网抗诉
	Restags *string `json:"restags,omitempty"`
}

func (o VmsTemplateStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VmsTemplateStatus struct{}"
	}

	return strings.Join([]string{"VmsTemplateStatus", string(data)}, " ")
}
