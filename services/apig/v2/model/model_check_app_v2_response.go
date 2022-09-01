package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CheckAppV2Response struct {

	// 编号
	Id *string `json:"id,omitempty" xml:"id"`

	// 名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 描述
	Remark         *string `json:"remark,omitempty" xml:"remark"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckAppV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckAppV2Response struct{}"
	}

	return strings.Join([]string{"CheckAppV2Response", string(data)}, " ")
}
