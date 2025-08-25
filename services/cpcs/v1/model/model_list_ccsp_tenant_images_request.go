package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCcspTenantImagesRequest Request Object
type ListCcspTenantImagesRequest struct {

	// 指定查询返回记录条数，默认值10
	PageSize *int32 `json:"page_size,omitempty"`

	// 索引位置，从page_num指定的下一条数据开始查询默认值为0
	PageNum *int32 `json:"page_num,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 镜像所属的服务类型： - **ENCRYPT_DECRYPT** : 加解密服务； - **SIGN_VERIFY** : 签名验签服务； - **KMS** : 密钥管理服务； - **TIMESTAMP** : 时间戳服务； - **COLLA_SIGN** : 协同签名服务； - **OTP** : 动态令牌服务； - **DB_ENCRYPT** : 数据库加密服务； - **FILE_ENCRYPT** : 文件加密服务 - **TIMESTAMP** : 时间戳服务； - **DIGIT_SEAL** : 电子签章服务； - **SSL_VPN** : ssl vpn服务；
	ServiceType *string `json:"service_type,omitempty"`

	// 排序属性，目前支持以下属性： - **create_time** : 镜像的创建时间（默认）
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方向，支持以下值： - **DESC** : 降序（默认） - **ASC** : 升序
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListCcspTenantImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCcspTenantImagesRequest struct{}"
	}

	return strings.Join([]string{"ListCcspTenantImagesRequest", string(data)}, " ")
}
