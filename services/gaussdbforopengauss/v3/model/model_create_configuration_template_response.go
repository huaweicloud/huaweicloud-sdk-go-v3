package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateConfigurationTemplateResponse struct {

	// 参数模板ID。
	Id *string `json:"id,omitempty"`

	// 参数模板名称。
	Name *string `json:"name,omitempty"`

	// 创建时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	CreatedAt      *string `json:"created_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateConfigurationTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateConfigurationTemplateResponse", string(data)}, " ")
}
