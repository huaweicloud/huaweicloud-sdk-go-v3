package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowResourceCopiesSummaryResponse Response Object
type ShowResourceCopiesSummaryResponse struct {

	// 企业项目ID
	EpsId *string `json:"eps_id,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 资源类型：Server-云服务器，Volume-云硬盘，Sfs-Turbo-高性能文件存储，Workspace-云桌面，MySQL-MySQL，PostgreSQL-PostgreSQL，SQLServer-SQLServer，MariaDB-MariaDB，GaussDB-GaussDB
	ResourceType *ShowResourceCopiesSummaryResponseResourceType `json:"resource_type,omitempty"`

	// 副本类型
	CopyType *string `json:"copy_type,omitempty"`

	// 总副本数
	TotalCopyCounts *int32 `json:"total_copy_counts,omitempty"`

	// 完成的备份总数, key -> ResourceCopyStatisticsKeyEnum.COMPLETED.getValue()
	CompletedCopyCounts *int32 `json:"completed_copy_counts,omitempty"`

	// 失败的备份总数, key -> ResourceCopyStatisticsKeyEnum.FAILED.getValue()
	FailedCopyCounts *int32 `json:"failed_copy_counts,omitempty"`

	// 统计信息
	Summary        *[]GetResourceCopySummaryResponseSummary `json:"summary,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o ShowResourceCopiesSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceCopiesSummaryResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceCopiesSummaryResponse", string(data)}, " ")
}

type ShowResourceCopiesSummaryResponseResourceType struct {
	value string
}

type ShowResourceCopiesSummaryResponseResourceTypeEnum struct {
	SERVER      ShowResourceCopiesSummaryResponseResourceType
	VOLUME      ShowResourceCopiesSummaryResponseResourceType
	SFS_TURBO   ShowResourceCopiesSummaryResponseResourceType
	WORKSPACE   ShowResourceCopiesSummaryResponseResourceType
	MY_SQL      ShowResourceCopiesSummaryResponseResourceType
	POSTGRE_SQL ShowResourceCopiesSummaryResponseResourceType
	SQL_SERVER  ShowResourceCopiesSummaryResponseResourceType
	MARIA_DB    ShowResourceCopiesSummaryResponseResourceType
	GAUSS_DB    ShowResourceCopiesSummaryResponseResourceType
}

func GetShowResourceCopiesSummaryResponseResourceTypeEnum() ShowResourceCopiesSummaryResponseResourceTypeEnum {
	return ShowResourceCopiesSummaryResponseResourceTypeEnum{
		SERVER: ShowResourceCopiesSummaryResponseResourceType{
			value: "Server",
		},
		VOLUME: ShowResourceCopiesSummaryResponseResourceType{
			value: "Volume",
		},
		SFS_TURBO: ShowResourceCopiesSummaryResponseResourceType{
			value: "Sfs-Turbo",
		},
		WORKSPACE: ShowResourceCopiesSummaryResponseResourceType{
			value: "Workspace",
		},
		MY_SQL: ShowResourceCopiesSummaryResponseResourceType{
			value: "MySQL",
		},
		POSTGRE_SQL: ShowResourceCopiesSummaryResponseResourceType{
			value: "PostgreSQL",
		},
		SQL_SERVER: ShowResourceCopiesSummaryResponseResourceType{
			value: "SQLServer",
		},
		MARIA_DB: ShowResourceCopiesSummaryResponseResourceType{
			value: "MariaDB",
		},
		GAUSS_DB: ShowResourceCopiesSummaryResponseResourceType{
			value: "GaussDB",
		},
	}
}

func (c ShowResourceCopiesSummaryResponseResourceType) Value() string {
	return c.value
}

func (c ShowResourceCopiesSummaryResponseResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResourceCopiesSummaryResponseResourceType) UnmarshalJSON(b []byte) error {
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
