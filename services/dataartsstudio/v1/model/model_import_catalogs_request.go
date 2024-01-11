package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportCatalogsRequest Request Object
type ImportCatalogsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 需要执行的动作
	ActionId string `json:"action-id"`

	// 是否需要覆盖更新已有的主题
	SkipExist *bool `json:"skip-exist,omitempty"`

	Body *ImportCatalogsRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportCatalogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportCatalogsRequest struct{}"
	}

	return strings.Join([]string{"ImportCatalogsRequest", string(data)}, " ")
}
