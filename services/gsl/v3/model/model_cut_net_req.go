package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CutNetReq struct {

	// 操作类型(ADD：断网，DEL:取消断网)
	Action string `json:"action"`
}

func (o CutNetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CutNetReq struct{}"
	}

	return strings.Join([]string{"CutNetReq", string(data)}, " ")
}
