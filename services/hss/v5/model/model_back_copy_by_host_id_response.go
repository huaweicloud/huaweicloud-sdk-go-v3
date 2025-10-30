package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackCopyByHostIdResponse struct {

	// **参数解释** 备份ID **取值范围** 字符长度0-65535位
	BackupId *string `json:"backup_id,omitempty"`

	// **参数解释** 备份名称 **取值范围** 字符长度0-65535位
	BackupName *string `json:"backup_name,omitempty"`

	// **参数解释** 备份状态 **取值范围** 字符长度0-65535位
	BackupStatus *string `json:"backup_status,omitempty"`

	// **参数解释** 创建时间 **取值范围** 取值0-2147483647
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释** 备份注册镜像ID列表 **取值范围** 取值0-200
	OsImagesData *[]ImageData `json:"os_images_data,omitempty"`

	// **参数解释** 备份标识 **取值范围** - 0：定时周期 - 1：勒索加密
	BackupTag *int32 `json:"backup_tag,omitempty"`
}

func (o BackCopyByHostIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackCopyByHostIdResponse struct{}"
	}

	return strings.Join([]string{"BackCopyByHostIdResponse", string(data)}, " ")
}
