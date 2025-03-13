package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PermissionApprovalDetailDto 审批详细信息
type PermissionApprovalDetailDto struct {

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 数据源类型 - HIVE数据源 - DWS数据源 - [DLI数据源](tag:nohcs)
	DatasourceType *PermissionApprovalDetailDtoDatasourceType `json:"datasource_type,omitempty"`

	// 超时时间
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// 申请权限详情列表
	Permissions *[]PermissionApprovalDetailDtoPermissions `json:"permissions,omitempty"`

	// 申请人详情列表
	Proposers *[]PermissionApprovalDetailDtoProposers `json:"proposers,omitempty"`
}

func (o PermissionApprovalDetailDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionApprovalDetailDto struct{}"
	}

	return strings.Join([]string{"PermissionApprovalDetailDto", string(data)}, " ")
}

type PermissionApprovalDetailDtoDatasourceType struct {
	value string
}

type PermissionApprovalDetailDtoDatasourceTypeEnum struct {
	HIVE PermissionApprovalDetailDtoDatasourceType
	DWS  PermissionApprovalDetailDtoDatasourceType
	DLI  PermissionApprovalDetailDtoDatasourceType
}

func GetPermissionApprovalDetailDtoDatasourceTypeEnum() PermissionApprovalDetailDtoDatasourceTypeEnum {
	return PermissionApprovalDetailDtoDatasourceTypeEnum{
		HIVE: PermissionApprovalDetailDtoDatasourceType{
			value: "HIVE",
		},
		DWS: PermissionApprovalDetailDtoDatasourceType{
			value: "DWS",
		},
		DLI: PermissionApprovalDetailDtoDatasourceType{
			value: "DLI",
		},
	}
}

func (c PermissionApprovalDetailDtoDatasourceType) Value() string {
	return c.value
}

func (c PermissionApprovalDetailDtoDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionApprovalDetailDtoDatasourceType) UnmarshalJSON(b []byte) error {
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
