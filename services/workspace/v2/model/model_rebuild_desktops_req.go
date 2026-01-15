package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebuildDesktopsReq 重建系统盘的请求。
type RebuildDesktopsReq struct {

	// 计算机id列表。
	DesktopIds []string `json:"desktop_ids"`

	// 镜像类型 - 公共镜像：gold - 私有镜像：private - 市场镜像：market
	ImageType string `json:"image_type"`

	// 模板ID。
	ImageId string `json:"image_id"`

	EncryptType *EncryptType `json:"encrypt_type,omitempty"`

	// 密钥ID，encrypt_type为ENCRYPTED时必传。
	KmsId *string `json:"kms_id,omitempty"`

	// os类型。
	OsType *string `json:"os_type,omitempty"`

	// 立即重建时给用户预留的保存数据的时间（单位：分钟）。
	DelayTime *int32 `json:"delay_time,omitempty"`

	// 下发重建系统盘任务时，给用户发送的提示信息。
	Message *string `json:"message,omitempty"`

	// 订单ID，包周期桌面重建系统盘，涉及收费镜像时需传。
	OrderId *string `json:"order_id,omitempty"`

	// 企业项目ID，默认\"0。\"
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 处理类型 - ONLY_FOR_EXPAND：仅对新扩容桌面生效 - FOR_EXPAND_AND_IDLE：对新扩容桌面与空闲桌面生效 - FOR_EXPAND_AND_ALL：对新扩容桌面与已有全部桌面生效
	HandleType *string `json:"handle_type,omitempty"`
}

func (o RebuildDesktopsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebuildDesktopsReq struct{}"
	}

	return strings.Join([]string{"RebuildDesktopsReq", string(data)}, " ")
}
