package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListVulHandleHistoryRequest Request Object
type ListVulHandleHistoryRequest struct {

	// 企业项目ID
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 漏洞状态，包含如下：   - vul_status_unfix：未处理   - vul_status_ignored：已忽略   - vul_status_verified：验证中   - vul_status_fixing：修复中   - vul_status_fixed：修复成功   - vul_status_reboot：修复成功待重启   - vul_status_failed：修复失败   - vul_status_fix_after_reboot：请重启主机再次修复
	Status *[]string `json:"status,omitempty"`

	// 漏洞ID
	VulId *string `json:"vul_id,omitempty"`

	// 漏洞类型，包含如下:   - linux_vul：Linux漏洞   - windows_vul：Windows漏洞   - web_cms：Web-CMS漏洞   - app_vul：应用漏洞   - urgent_vul：应急漏洞
	VulType *string `json:"vul_type,omitempty"`

	// 资产重要性，包含如下:   - important：重要资产   - common：一般资产   - test：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// 服务器组
	GroupName *string `json:"group_name,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器IP
	HostIp *string `json:"host_ip,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 排序字段，包含如下：   - handle_time：处置时间
	SortKey *ListVulHandleHistoryRequestSortKey `json:"sort_key,omitempty"`

	// 排序顺序，若sort_key不为空，设置返回结果按照sort_key升序或降序排序，默认降序排序，包含如下：   - asc：升序   - desc：降序
	SortDir *ListVulHandleHistoryRequestSortDir `json:"sort_dir,omitempty"`

	// 服务器公网IP
	PublicIp *string `json:"public_ip,omitempty"`

	// 服务器私网IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 每页个数
	Limit int32 `json:"limit"`

	// 偏移量
	Offset int32 `json:"offset"`
}

func (o ListVulHandleHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulHandleHistoryRequest struct{}"
	}

	return strings.Join([]string{"ListVulHandleHistoryRequest", string(data)}, " ")
}

type ListVulHandleHistoryRequestSortKey struct {
	value string
}

type ListVulHandleHistoryRequestSortKeyEnum struct {
	HANDLE_TIME ListVulHandleHistoryRequestSortKey
}

func GetListVulHandleHistoryRequestSortKeyEnum() ListVulHandleHistoryRequestSortKeyEnum {
	return ListVulHandleHistoryRequestSortKeyEnum{
		HANDLE_TIME: ListVulHandleHistoryRequestSortKey{
			value: "handle_time",
		},
	}
}

func (c ListVulHandleHistoryRequestSortKey) Value() string {
	return c.value
}

func (c ListVulHandleHistoryRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVulHandleHistoryRequestSortKey) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListVulHandleHistoryRequestSortDir struct {
	value string
}

type ListVulHandleHistoryRequestSortDirEnum struct {
	ASC  ListVulHandleHistoryRequestSortDir
	DESC ListVulHandleHistoryRequestSortDir
}

func GetListVulHandleHistoryRequestSortDirEnum() ListVulHandleHistoryRequestSortDirEnum {
	return ListVulHandleHistoryRequestSortDirEnum{
		ASC: ListVulHandleHistoryRequestSortDir{
			value: "asc",
		},
		DESC: ListVulHandleHistoryRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListVulHandleHistoryRequestSortDir) Value() string {
	return c.value
}

func (c ListVulHandleHistoryRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVulHandleHistoryRequestSortDir) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
