package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2NetworkBriefInfo 网络配置简要信息（列表返回）。
type V2NetworkBriefInfo struct {

	// **参数解释**：网络连接的ID。 **取值范围**：长度为1~64的英文、数字、中划线组合。
	Id *string `json:"id,omitempty"`

	// **参数解释**：路由网段范围，标准的IPv4格式。 **取值范围**：允许网段为10.0.0.0~10.255.0.0/8~22,172.16.0.0~172.31.0.0/12~22,192.168.0.0~192.168.0.0/16~22。
	CidrInVpc *string `json:"cidr_in_vpc,omitempty"`

	// **参数解释**：资源创建时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
}

func (o V2NetworkBriefInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2NetworkBriefInfo struct{}"
	}

	return strings.Join([]string{"V2NetworkBriefInfo", string(data)}, " ")
}
