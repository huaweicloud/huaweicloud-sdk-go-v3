package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MasterEipRequestSpecSpec struct {

	// 弹性网卡ID，必选参数
	Id *string `json:"id,omitempty"`
}

func (o MasterEipRequestSpecSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterEipRequestSpecSpec struct{}"
	}

	return strings.Join([]string{"MasterEipRequestSpecSpec", string(data)}, " ")
}
