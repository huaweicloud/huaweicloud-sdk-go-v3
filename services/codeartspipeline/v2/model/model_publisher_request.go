package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublisherRequest struct {

	// 名称
	Name string `json:"name"`

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 图标URL
	LogoUrl *string `json:"logo_url,omitempty"`

	// 网页地址
	Website *string `json:"website,omitempty"`

	// 地址
	SupportUrl string `json:"support_url"`

	// 地址
	SourceUrl *string `json:"source_url,omitempty"`

	// 英文名
	EnName string `json:"en_name"`

	// 唯一ID
	PublisherUniqueId *string `json:"publisher_unique_id,omitempty"`
}

func (o PublisherRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublisherRequest struct{}"
	}

	return strings.Join([]string{"PublisherRequest", string(data)}, " ")
}
