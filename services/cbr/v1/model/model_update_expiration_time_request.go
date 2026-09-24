package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateExpirationTimeRequest Request Object
type UpdateExpirationTimeRequest struct {

	// 存储库ID，默认取值不涉及。 [获取方法请参见\"[获取存储库ID](https://support.huaweicloud.com/api-cbr/ListVault.html)\"。](tag:hws) [获取方法请参见\"[获取存储库ID](https://support.huaweicloud.com/intl/zh-cn/api-cbr/ListVault.html)\"。](tag:hws_hk)
	VaultId string `json:"vault_id"`

	Body *UpdateExpirationTimeReq `json:"body,omitempty"`
}

func (o UpdateExpirationTimeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExpirationTimeRequest struct{}"
	}

	return strings.Join([]string{"UpdateExpirationTimeRequest", string(data)}, " ")
}
