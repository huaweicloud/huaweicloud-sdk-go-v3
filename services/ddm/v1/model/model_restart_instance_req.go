package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartInstanceReq This is a auto restart Body Object
type RestartInstanceReq struct {
	Restart *RestarInstanceInfo `json:"restart,omitempty"`
}

func (o RestartInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartInstanceReq struct{}"
	}

	return strings.Join([]string{"RestartInstanceReq", string(data)}, " ")
}
