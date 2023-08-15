package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGaussMySqlInstanceDetailInfoRequest Request Object
type ListGaussMySqlInstanceDetailInfoRequest struct {

	// 语言。默认英语。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。最多同时输入20个实例ID，用英文逗号分隔。
	InstanceIds string `json:"instance_ids"`
}

func (o ListGaussMySqlInstanceDetailInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussMySqlInstanceDetailInfoRequest struct{}"
	}

	return strings.Join([]string{"ListGaussMySqlInstanceDetailInfoRequest", string(data)}, " ")
}
