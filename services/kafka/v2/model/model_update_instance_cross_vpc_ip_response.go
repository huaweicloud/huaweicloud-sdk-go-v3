package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceCrossVpcIpResponse Response Object
type UpdateInstanceCrossVpcIpResponse struct {

	// **参数解释**： 修改跨VPC访问结果。 **取值范围**： - true：修改跨VPC访问成功。 - false：修改跨VPC访问失败。
	Success *bool `json:"success,omitempty"`

	// **参数解释**： 修改broker跨VPC访问的结果列表。
	Results        *[]UpdateInstanceCrossVpcIpRespResults `json:"results,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o UpdateInstanceCrossVpcIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceCrossVpcIpResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceCrossVpcIpResponse", string(data)}, " ")
}
