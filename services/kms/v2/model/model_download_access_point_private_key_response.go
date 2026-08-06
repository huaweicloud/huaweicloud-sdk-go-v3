package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAccessPointPrivateKeyResponse Response Object
type DownloadAccessPointPrivateKeyResponse struct {

	// **参数解释：** 接入点ID **取值范围：** 不涉及
	AccessPointId *string `json:"access_point_id,omitempty"`

	// **参数解释：** 通用类型接入点私钥 **取值范围：** 不涉及
	PrivateKey     *string `json:"private_key,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DownloadAccessPointPrivateKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAccessPointPrivateKeyResponse struct{}"
	}

	return strings.Join([]string{"DownloadAccessPointPrivateKeyResponse", string(data)}, " ")
}
