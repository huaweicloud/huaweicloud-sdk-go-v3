package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Connection struct {

	// - 参数解释：下联面网口主备IP；ESW实例在本端二层子网中占用的主备接口IP。 - 约束限制：字符串列表，只能设置两个字符串，且每个字符串内容应该是标准IPv4格式；IP必须在二层子网网段范围内。 - 取值范围：不涉及。 - 默认取值：不涉及。
	FixedIps []string `json:"fixed_ips"`

	// - 参数解释：二层连接的唯一标识。 - 约束限制：带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Id string `json:"id"`

	// - 参数解释：ESW实例的唯一标识。 - 约束限制：带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	InstanceId string `json:"instance_id"`

	// - 参数解释：二层连接的名称。 - 约束限制：   - 长度范围为1~64个字符。   - 名称由中文、英文字母、数字、下划线（_）、中划线（-）、点（.）组成。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Name string `json:"name"`

	// - 参数解释：ESW实例所属项目ID。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	ProjectId string `json:"project_id"`

	// - 参数解释：远端隧道信息。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	RemoteInfos []RemoteInfosResult `json:"remote_infos"`

	// - 参数解释：二层连接的状态。 - 约束限制：不涉及。 - 取值范围：   - pending：准备中   - connected：已连接   - disconnect：未连接   - failed：创建失败   - abnormal：异常 - 默认取值：不涉及。
	Status string `json:"status"`

	// - 参数解释：二层连接关联的二层子网ID。 - 约束限制：   - 需要使用本租户下可操作的子网资源的ID；此值即为子网详情中的“网络ID”参数值。   - 带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	VirsubnetId string `json:"virsubnet_id"`

	// - 参数解释：ESW所在VPC资源ID。 - 约束限制：   - 需要使用本租户下可操作的VPC资源的ID。   - 带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	VpcId string `json:"vpc_id"`

	// - 参数解释：二层连接创建时间。 - 约束限制：UTC时间，格式为yyyy-MM-ddTHH:mm:ss。 - 取值范围：不涉及。 - 默认取值：不涉及。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// - 参数解释：二层连接更新时间。 - 约束限制：UTC时间，格式为yyyy-MM-ddTHH:mm:ss。 - 取值范围：不涉及。 - 默认取值：不涉及。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o Connection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Connection struct{}"
	}

	return strings.Join([]string{"Connection", string(data)}, " ")
}
