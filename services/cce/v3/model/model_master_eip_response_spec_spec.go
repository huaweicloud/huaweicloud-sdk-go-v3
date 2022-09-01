package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MasterEipResponseSpecSpec struct {

	// 弹性网卡ID
	Id *string `json:"id,omitempty" xml:"id"`

	Eip *EipSpec `json:"eip,omitempty" xml:"eip"`

	// 是否动态创建
	IsDynamic *bool `json:"IsDynamic,omitempty" xml:"IsDynamic"`
}

func (o MasterEipResponseSpecSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterEipResponseSpecSpec struct{}"
	}

	return strings.Join([]string{"MasterEipResponseSpecSpec", string(data)}, " ")
}
