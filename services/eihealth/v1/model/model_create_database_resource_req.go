package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDatabaseResourceReq struct {

	// 规格编码
	SpecCode string `json:"spec_code"`

	// 磁盘存储空间，该字段暂不生效
	DiskSpace *int32 `json:"disk_space,omitempty"`

	// 磁盘是否加密
	DiskEncrypt bool `json:"disk_encrypt"`
}

func (o CreateDatabaseResourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseResourceReq struct{}"
	}

	return strings.Join([]string{"CreateDatabaseResourceReq", string(data)}, " ")
}
