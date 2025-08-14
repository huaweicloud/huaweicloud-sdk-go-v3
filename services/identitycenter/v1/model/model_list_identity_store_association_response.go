package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIdentityStoreAssociationResponse Response Object
type ListIdentityStoreAssociationResponse struct {

	// IAM身份中心服务实例关联的身份源配置信息
	IdentityStoreAssociations *[]IdentityStoreDto `json:"identity_store_associations,omitempty"`
	HttpStatusCode            int                 `json:"-"`
}

func (o ListIdentityStoreAssociationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIdentityStoreAssociationResponse struct{}"
	}

	return strings.Join([]string{"ListIdentityStoreAssociationResponse", string(data)}, " ")
}
