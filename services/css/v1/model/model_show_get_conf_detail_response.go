package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowGetConfDetailResponse struct {

	// 配置文件名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 配置文件状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 配置文件内容。
	ConfContent *string `json:"confContent,omitempty" xml:"confContent"`

	Setting *Setting `json:"setting,omitempty" xml:"setting"`

	// 更新时间。
	UpdateAt       *string `json:"updateAt,omitempty" xml:"updateAt"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowGetConfDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGetConfDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowGetConfDetailResponse", string(data)}, " ")
}
