package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// notebook详情
type NotebookEntity struct {

	// notebook ID
	Id *string `json:"id,omitempty"`

	// notebook名称
	Name *string `json:"name,omitempty"`

	// notebook描述
	Description *string `json:"description,omitempty"`

	// notebook所属用户
	Creator *string `json:"creator,omitempty"`

	// notebook访问URL
	Url *string `json:"url,omitempty"`

	Flavor *FlavorInfo `json:"flavor,omitempty"`

	Status *NotebookStatus `json:"status,omitempty"`

	Image *NotebookImage `json:"image,omitempty"`

	// notebook存储信息
	Storages *[]NotebookStorage `json:"storages,omitempty"`

	// notebook创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// notebook更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// notebook失败信息
	FailedMessage *string `json:"failed_message,omitempty"`

	// cce事件
	Events *[]TaskEventRsp `json:"events,omitempty"`
}

func (o NotebookEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotebookEntity struct{}"
	}

	return strings.Join([]string{"NotebookEntity", string(data)}, " ")
}
