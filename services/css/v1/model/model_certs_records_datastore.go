package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CertsRecordsDatastore 证书记录。
type CertsRecordsDatastore struct {

	// 证书记录ID。
	Id *string `json:"id,omitempty"`

	// 证书状态。
	Status *string `json:"status,omitempty"`

	// 证书记录文件位置。
	FileLocation *string `json:"fileLocation,omitempty"`

	// 证书记录文件名称。
	FileName *bool `json:"fileName,omitempty"`

	// 证书记录更新时间。
	UpdateAt *bool `json:"updateAt,omitempty"`
}

func (o CertsRecordsDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertsRecordsDatastore struct{}"
	}

	return strings.Join([]string{"CertsRecordsDatastore", string(data)}, " ")
}
