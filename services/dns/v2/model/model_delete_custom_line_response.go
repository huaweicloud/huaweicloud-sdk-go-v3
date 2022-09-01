package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteCustomLineResponse struct {

	// 解析线路ID。
	LineId *string `json:"line_id,omitempty" xml:"line_id"`

	// 解析线路名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// IP地址段。
	IpSegments *[]string `json:"ip_segments,omitempty" xml:"ip_segments"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at"`

	// 资源状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 自定义线路的描述信息。
	Description    *string `json:"description,omitempty" xml:"description"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteCustomLineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomLineResponse struct{}"
	}

	return strings.Join([]string{"DeleteCustomLineResponse", string(data)}, " ")
}
