package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyDomainTemplateResponse Response Object
type ApplyDomainTemplateResponse struct {

	// **参数解释：** 操作ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释：** 应用模板状态（任务粒度） **约束限制：** 不涉及 **取值范围：** - success: 应用模板成功 - fail: 应用模板失败 **默认取值：** 不涉及
	Status *string `json:"status,omitempty"`

	Detail         *[]ApplyTmlDetail `json:"detail,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ApplyDomainTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyDomainTemplateResponse struct{}"
	}

	return strings.Join([]string{"ApplyDomainTemplateResponse", string(data)}, " ")
}
