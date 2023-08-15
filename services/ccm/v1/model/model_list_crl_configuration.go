package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListCrlConfiguration struct {

	// 是否启用CRL发布功能。 - **true** - **false**
	Enabled bool `json:"enabled"`

	// 吊销列表文件名称。 > 若用户不指定，系统将默认采用当前证书的父CA ID。
	CrlName string `json:"crl_name"`

	// OBS桶名称。
	ObsBucketName string `json:"obs_bucket_name"`

	// CRL更新周期，单位为\"天\"。当启用CRL发布功能，为必填项。
	ValidDays int32 `json:"valid_days"`

	// 吊销列表分发地址，即对应的OBS桶中的CRL文件地址。 > 本参数由程序根据crl_name、obs_bucket_name以及OBS地址进行拼接而成。
	CrlDisPoint string `json:"crl_dis_point"`
}

func (o ListCrlConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCrlConfiguration struct{}"
	}

	return strings.Join([]string{"ListCrlConfiguration", string(data)}, " ")
}
