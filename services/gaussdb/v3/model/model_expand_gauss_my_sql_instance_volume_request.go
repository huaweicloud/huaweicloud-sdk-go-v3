package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExpandGaussMySqlInstanceVolumeRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *MysqlExtendInstanceVolumeRequest `json:"body,omitempty" xml:"body"`
}

func (o ExpandGaussMySqlInstanceVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandGaussMySqlInstanceVolumeRequest struct{}"
	}

	return strings.Join([]string{"ExpandGaussMySqlInstanceVolumeRequest", string(data)}, " ")
}
