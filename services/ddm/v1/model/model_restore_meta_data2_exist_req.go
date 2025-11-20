package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreMetaData2ExistReq struct {

	// metadata恢复源。
	Source *interface{} `json:"source,omitempty"`

	// metadata恢复目标。
	Target *interface{} `json:"target,omitempty"`
}

func (o RestoreMetaData2ExistReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreMetaData2ExistReq struct{}"
	}

	return strings.Join([]string{"RestoreMetaData2ExistReq", string(data)}, " ")
}
