package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInferApiKeyResponse Response Object
type DeleteInferApiKeyResponse struct {

	// **参数解释：** api-key的ID，在[创建API_KEY](CreateInferApiKey.xml)时即可在返回体中获取，也可通过[查询api-keys列表](ListInferApiKeys.xml)获取当前用户拥有的api-key，其中id字段即为api-key的ID。 **取值范围：** UUID格式。
	KeyId *string `json:"key_id,omitempty"`

	// **参数解释：** api-key的名称，在[创建API_KEY](CreateInferApiKey.xml)时自定义。 **取值范围：** 支持1-64个字符，可以包含字母、汉字、数字、连字符和下划线。
	Name *string `json:"name,omitempty"`

	// **参数解释：** api-key的描述，在[创建API_KEY](CreateInferApiKey.xml)时自定义。 **取值范围：** 支持1-256个字符，可以包含字母、汉字、数字、连字符和下划线。
	Description *string `json:"description,omitempty"`

	// **参数解释：** api-key的创建时间，根据创建时的当前时间自动生成。 **取值范围：** 毫秒级时间戳，13位数字，如1609459200000。
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释：** api-key生效范围。 **取值范围：** - USER：表示生效范围为用户级别，可以访问该用户创建的所有在线服务。 - SERVICE：表示生效范围为单个服务，可以访问绑定该api-key的在线服务。
	Scope *string `json:"scope,omitempty"`

	// **参数解释：** 用户domain ID。获取方法请参见[获取账号名和账号ID](modelarts_03_0148.xml)。 **取值范围：** 账号ID。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释：** [用户项目ID](tag:hws,hws_hk,fcs,fcs_super)[资源空间ID](tag:hcs,hcs_sm)。获取方法请参见[[获取项目ID和名称](tag:hws,hws_hk,fcs,fcs_super)[获取资源空间ID和名称](tag:hcs,hcs_sm)](modelarts_03_0147.xml)。 **取值范围：** 账号的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 绑定此api-key的在线服务列表。
	Services *[]ServiceIdName `json:"services,omitempty"`

	// **参数解释：** 工作空间ID。 **取值范围：** 工作空间ID。
	WorkspaceId    *string `json:"workspace_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInferApiKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInferApiKeyResponse struct{}"
	}

	return strings.Join([]string{"DeleteInferApiKeyResponse", string(data)}, " ")
}
