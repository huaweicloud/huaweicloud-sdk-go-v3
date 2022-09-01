package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 密钥库详情
type KeystoreDetails struct {

	// 密钥库ID
	KeystoreId *string `json:"keystore_id,omitempty" xml:"keystore_id"`

	// 用户域ID
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 密钥库别名
	KeystoreAlias *string `json:"keystore_alias,omitempty" xml:"keystore_alias"`

	// 密钥库类型
	KeystoreType *string `json:"keystore_type,omitempty" xml:"keystore_type"`

	// DHSM集群id，要求集群当前未创建专属密钥库
	HsmClusterId *string `json:"hsm_cluster_id,omitempty" xml:"hsm_cluster_id"`

	// 密钥库创建时间，UTC时间戳。
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`
}

func (o KeystoreDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoreDetails struct{}"
	}

	return strings.Join([]string{"KeystoreDetails", string(data)}, " ")
}
