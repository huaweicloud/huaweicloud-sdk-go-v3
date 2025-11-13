package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreatePublicRecordsetsTaskItem struct {

	// **参数解释：** 托管该记录的域名。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneName string `json:"zone_name"`

	// **参数解释：** 主机记录或域名。 末尾有点则认为是完整域名，如果不包含点，则认为是域名前缀。例如： 输入\"www\"，创建成功后域名显示为\"www.example.com.\" 输入\"www.example.com.\"（注意域名末尾有点），创建成功后域名显示为\"www.example.com.\" 输入\"www.example.com\"（注意域名末尾没有点），创建成功后域名显示为\"www.example.com.example.com.\" **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 记录集的类型。 **约束限制：** 不涉及。 **取值范围：** - 公网域名的记录类型: A、CNAME、MX、AAAA、TXT、SRV、NS、CAA。  **默认取值：** 不涉及。
	Type string `json:"type"`

	// **参数解释：** 解析线路ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** default_view。
	Line *string `json:"line,omitempty"`

	// **参数解释：** 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。 **约束限制：** 不涉及。 **取值范围：** 1~2147483647。 **默认取值：** 300
	Ttl *int32 `json:"ttl,omitempty"`

	// **参数解释：** 解析记录的权重。 **约束限制：** MX记录集不支持权重，创建时会忽略权重字段，如果您的MX记录集有多个值，在一条记录集里写多个值即可。(若存在多条同主机记录、同线路类型的MX记录，会自动合并为一条创建) **取值范围：** 0~1000。 **默认取值：** 不涉及。
	Weight *int32 `json:"weight,omitempty"`

	// **参数解释：** 解析记录的值。不同类型解析记录对应的值的规则不同。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Records []string `json:"records"`

	// **参数解释：** 解析记录的状态。 **约束限制：** 不涉及。 **取值范围：** - ENABLE：启用解析 - DISABLE：暂停解析  **默认取值：** ENABLE。
	Status *string `json:"status,omitempty"`
}

func (o BatchCreatePublicRecordsetsTaskItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePublicRecordsetsTaskItem struct{}"
	}

	return strings.Join([]string{"BatchCreatePublicRecordsetsTaskItem", string(data)}, " ")
}
