package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 克隆服务器类
type CloneServer struct {

	// 克隆服务器ID
	VmId *string `json:"vm_id,omitempty" xml:"vm_id"`

	// 克隆虚拟机的名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 克隆错误信息
	CloneError *string `json:"clone_error,omitempty" xml:"clone_error"`

	// 克隆状态
	CloneState *string `json:"clone_state,omitempty" xml:"clone_state"`

	// 克隆错误信息描述
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`
}

func (o CloneServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloneServer struct{}"
	}

	return strings.Join([]string{"CloneServer", string(data)}, " ")
}
