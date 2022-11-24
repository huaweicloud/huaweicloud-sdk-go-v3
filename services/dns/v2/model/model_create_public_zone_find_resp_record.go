package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePublicZoneFindRespRecord struct {

	// 找回记录host名称。
	Host *string `json:"host,omitempty"`

	// 找回记录解析值。
	Value *string `json:"value,omitempty"`
}

func (o CreatePublicZoneFindRespRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicZoneFindRespRecord struct{}"
	}

	return strings.Join([]string{"CreatePublicZoneFindRespRecord", string(data)}, " ")
}
