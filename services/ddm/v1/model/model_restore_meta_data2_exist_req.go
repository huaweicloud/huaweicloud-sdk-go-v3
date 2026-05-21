package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreMetaData2ExistReq struct {
	Source *RestoreMetaDataSource `json:"source"`

	Target *RestoreMetaDataTarget `json:"target"`
}

func (o RestoreMetaData2ExistReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreMetaData2ExistReq struct{}"
	}

	return strings.Join([]string{"RestoreMetaData2ExistReq", string(data)}, " ")
}
