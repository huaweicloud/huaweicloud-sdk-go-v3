package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ec/v1/model"
)

type EcClient struct {
	HcClient *http_client.HcHttpClient
}

func NewEcClient(hcClient *http_client.HcHttpClient) *EcClient {
	return &EcClient{HcClient: hcClient}
}

func EcClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// CreateEcnAccessPoint 添加新的接入点
//
// 添加新的接入点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) CreateEcnAccessPoint(request *model.CreateEcnAccessPointRequest) (*model.CreateEcnAccessPointResponse, error) {
	requestDef := GenReqDefForCreateEcnAccessPoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEcnAccessPointResponse), nil
	}
}

// CreateEcnAccessPointInvoker 添加新的接入点
func (c *EcClient) CreateEcnAccessPointInvoker(request *model.CreateEcnAccessPointRequest) *CreateEcnAccessPointInvoker {
	requestDef := GenReqDefForCreateEcnAccessPoint()
	return &CreateEcnAccessPointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEcnAccessPoint 删除接入点
//
// 根据企业连接网络ID，删除接入点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) DeleteEcnAccessPoint(request *model.DeleteEcnAccessPointRequest) (*model.DeleteEcnAccessPointResponse, error) {
	requestDef := GenReqDefForDeleteEcnAccessPoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEcnAccessPointResponse), nil
	}
}

// DeleteEcnAccessPointInvoker 删除接入点
func (c *EcClient) DeleteEcnAccessPointInvoker(request *model.DeleteEcnAccessPointRequest) *DeleteEcnAccessPointInvoker {
	requestDef := GenReqDefForDeleteEcnAccessPoint()
	return &DeleteEcnAccessPointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEcnAccessPointByEcnId 查询接入点
//
// 根据企业连接网络ID，查询其下所有接入点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ListEcnAccessPointByEcnId(request *model.ListEcnAccessPointByEcnIdRequest) (*model.ListEcnAccessPointByEcnIdResponse, error) {
	requestDef := GenReqDefForListEcnAccessPointByEcnId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEcnAccessPointByEcnIdResponse), nil
	}
}

// ListEcnAccessPointByEcnIdInvoker 查询接入点
func (c *EcClient) ListEcnAccessPointByEcnIdInvoker(request *model.ListEcnAccessPointByEcnIdRequest) *ListEcnAccessPointByEcnIdInvoker {
	requestDef := GenReqDefForListEcnAccessPointByEcnId()
	return &ListEcnAccessPointByEcnIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEcnAccessPoint 更新接入点
//
// 根据企业连接网络ID，更新接入点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) UpdateEcnAccessPoint(request *model.UpdateEcnAccessPointRequest) (*model.UpdateEcnAccessPointResponse, error) {
	requestDef := GenReqDefForUpdateEcnAccessPoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEcnAccessPointResponse), nil
	}
}

// UpdateEcnAccessPointInvoker 更新接入点
func (c *EcClient) UpdateEcnAccessPointInvoker(request *model.UpdateEcnAccessPointRequest) *UpdateEcnAccessPointInvoker {
	requestDef := GenReqDefForUpdateEcnAccessPoint()
	return &UpdateEcnAccessPointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddEcnWithIeg 绑定智能企业网关到企业连接网络
//
// 绑定智能企业网关到企业连接网络
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) AddEcnWithIeg(request *model.AddEcnWithIegRequest) (*model.AddEcnWithIegResponse, error) {
	requestDef := GenReqDefForAddEcnWithIeg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddEcnWithIegResponse), nil
	}
}

// AddEcnWithIegInvoker 绑定智能企业网关到企业连接网络
func (c *EcClient) AddEcnWithIegInvoker(request *model.AddEcnWithIegRequest) *AddEcnWithIegInvoker {
	requestDef := GenReqDefForAddEcnWithIeg()
	return &AddEcnWithIegInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEcnWithIeg 解除智能企业网关和企业连接网络的绑定
//
// 解除智能企业网关和企业连接网络的绑定
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) DeleteEcnWithIeg(request *model.DeleteEcnWithIegRequest) (*model.DeleteEcnWithIegResponse, error) {
	requestDef := GenReqDefForDeleteEcnWithIeg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEcnWithIegResponse), nil
	}
}

// DeleteEcnWithIegInvoker 解除智能企业网关和企业连接网络的绑定
func (c *EcClient) DeleteEcnWithIegInvoker(request *model.DeleteEcnWithIegRequest) *DeleteEcnWithIegInvoker {
	requestDef := GenReqDefForDeleteEcnWithIeg()
	return &DeleteEcnWithIegInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEcn 查询企业连接网络列表
//
// 查询租户的企业连接网络列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ListEcn(request *model.ListEcnRequest) (*model.ListEcnResponse, error) {
	requestDef := GenReqDefForListEcn()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEcnResponse), nil
	}
}

// ListEcnInvoker 查询企业连接网络列表
func (c *EcClient) ListEcnInvoker(request *model.ListEcnRequest) *ListEcnInvoker {
	requestDef := GenReqDefForListEcn()
	return &ListEcnInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEcnWithIeg 查询企业连接网络与智能企业网关绑定关系
//
// 根据企业连接网络ID，查询企业连接网络与智能企业网关绑定关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ListEcnWithIeg(request *model.ListEcnWithIegRequest) (*model.ListEcnWithIegResponse, error) {
	requestDef := GenReqDefForListEcnWithIeg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEcnWithIegResponse), nil
	}
}

// ListEcnWithIegInvoker 查询企业连接网络与智能企业网关绑定关系
func (c *EcClient) ListEcnWithIegInvoker(request *model.ListEcnWithIegRequest) *ListEcnWithIegInvoker {
	requestDef := GenReqDefForListEcnWithIeg()
	return &ListEcnWithIegInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEcnInfo 查询企业连接网络
//
// 根据企业连接网络ID，查询企业连接网络
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ShowEcnInfo(request *model.ShowEcnInfoRequest) (*model.ShowEcnInfoResponse, error) {
	requestDef := GenReqDefForShowEcnInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEcnInfoResponse), nil
	}
}

// ShowEcnInfoInvoker 查询企业连接网络
func (c *EcClient) ShowEcnInfoInvoker(request *model.ShowEcnInfoRequest) *ShowEcnInfoInvoker {
	requestDef := GenReqDefForShowEcnInfo()
	return &ShowEcnInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEcnWithIeg 查询企业连接网络与单个智能企业网关绑定关系
//
// 查询企业连接网络与单个智能企业网关绑定关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ShowEcnWithIeg(request *model.ShowEcnWithIegRequest) (*model.ShowEcnWithIegResponse, error) {
	requestDef := GenReqDefForShowEcnWithIeg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEcnWithIegResponse), nil
	}
}

// ShowEcnWithIegInvoker 查询企业连接网络与单个智能企业网关绑定关系
func (c *EcClient) ShowEcnWithIegInvoker(request *model.ShowEcnWithIegRequest) *ShowEcnWithIegInvoker {
	requestDef := GenReqDefForShowEcnWithIeg()
	return &ShowEcnWithIegInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEcn 更新企业连接网络
//
// 根据企业连接网络ID，更新企业连接网络
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) UpdateEcn(request *model.UpdateEcnRequest) (*model.UpdateEcnResponse, error) {
	requestDef := GenReqDefForUpdateEcn()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEcnResponse), nil
	}
}

// UpdateEcnInvoker 更新企业连接网络
func (c *EcClient) UpdateEcnInvoker(request *model.UpdateEcnRequest) *UpdateEcnInvoker {
	requestDef := GenReqDefForUpdateEcn()
	return &UpdateEcnInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEquipment 激活智能企业网关设备
//
// 激活智能企业网关设备
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) CreateEquipment(request *model.CreateEquipmentRequest) (*model.CreateEquipmentResponse, error) {
	requestDef := GenReqDefForCreateEquipment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEquipmentResponse), nil
	}
}

// CreateEquipmentInvoker 激活智能企业网关设备
func (c *EcClient) CreateEquipmentInvoker(request *model.CreateEquipmentRequest) *CreateEquipmentInvoker {
	requestDef := GenReqDefForCreateEquipment()
	return &CreateEquipmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEquipment 移除智能企业网关设备
//
// 移除智能企业网关设备
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) DeleteEquipment(request *model.DeleteEquipmentRequest) (*model.DeleteEquipmentResponse, error) {
	requestDef := GenReqDefForDeleteEquipment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEquipmentResponse), nil
	}
}

// DeleteEquipmentInvoker 移除智能企业网关设备
func (c *EcClient) DeleteEquipmentInvoker(request *model.DeleteEquipmentRequest) *DeleteEquipmentInvoker {
	requestDef := GenReqDefForDeleteEquipment()
	return &DeleteEquipmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GenerateInitialConfiguration 生成智能企业网关设备初始配置
//
// 生成智能企业网关设备初始配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) GenerateInitialConfiguration(request *model.GenerateInitialConfigurationRequest) (*model.GenerateInitialConfigurationResponse, error) {
	requestDef := GenReqDefForGenerateInitialConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GenerateInitialConfigurationResponse), nil
	}
}

// GenerateInitialConfigurationInvoker 生成智能企业网关设备初始配置
func (c *EcClient) GenerateInitialConfigurationInvoker(request *model.GenerateInitialConfigurationRequest) *GenerateInitialConfigurationInvoker {
	requestDef := GenReqDefForGenerateInitialConfiguration()
	return &GenerateInitialConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEquipments 查询智能企业网关设备列表
//
// 根据智能企业网关ID，查询智能企业网关设备列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ListEquipments(request *model.ListEquipmentsRequest) (*model.ListEquipmentsResponse, error) {
	requestDef := GenReqDefForListEquipments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEquipmentsResponse), nil
	}
}

// ListEquipmentsInvoker 查询智能企业网关设备列表
func (c *EcClient) ListEquipmentsInvoker(request *model.ListEquipmentsRequest) *ListEquipmentsInvoker {
	requestDef := GenReqDefForListEquipments()
	return &ListEquipmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RebootEquipment 重启智能企业网关设备
//
// 重启智能企业网关设备
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) RebootEquipment(request *model.RebootEquipmentRequest) (*model.RebootEquipmentResponse, error) {
	requestDef := GenReqDefForRebootEquipment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RebootEquipmentResponse), nil
	}
}

// RebootEquipmentInvoker 重启智能企业网关设备
func (c *EcClient) RebootEquipmentInvoker(request *model.RebootEquipmentRequest) *RebootEquipmentInvoker {
	requestDef := GenReqDefForRebootEquipment()
	return &RebootEquipmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEquipmentInfo 查询智能企业网关设备
//
// 查询智能企业网关设备
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ShowEquipmentInfo(request *model.ShowEquipmentInfoRequest) (*model.ShowEquipmentInfoResponse, error) {
	requestDef := GenReqDefForShowEquipmentInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEquipmentInfoResponse), nil
	}
}

// ShowEquipmentInfoInvoker 查询智能企业网关设备
func (c *EcClient) ShowEquipmentInfoInvoker(request *model.ShowEquipmentInfoRequest) *ShowEquipmentInfoInvoker {
	requestDef := GenReqDefForShowEquipmentInfo()
	return &ShowEquipmentInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEquipmentSpecificConfig 查询智能企业网关设备基础规格配置
//
// 查询智能企业网关设备基础规格配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ShowEquipmentSpecificConfig(request *model.ShowEquipmentSpecificConfigRequest) (*model.ShowEquipmentSpecificConfigResponse, error) {
	requestDef := GenReqDefForShowEquipmentSpecificConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEquipmentSpecificConfigResponse), nil
	}
}

// ShowEquipmentSpecificConfigInvoker 查询智能企业网关设备基础规格配置
func (c *EcClient) ShowEquipmentSpecificConfigInvoker(request *model.ShowEquipmentSpecificConfigRequest) *ShowEquipmentSpecificConfigInvoker {
	requestDef := GenReqDefForShowEquipmentSpecificConfig()
	return &ShowEquipmentSpecificConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEquipmentEsn 修改智能企业网关设备ESN
//
// 修改智能企业网关设备ESN
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) UpdateEquipmentEsn(request *model.UpdateEquipmentEsnRequest) (*model.UpdateEquipmentEsnResponse, error) {
	requestDef := GenReqDefForUpdateEquipmentEsn()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEquipmentEsnResponse), nil
	}
}

// UpdateEquipmentEsnInvoker 修改智能企业网关设备ESN
func (c *EcClient) UpdateEquipmentEsnInvoker(request *model.UpdateEquipmentEsnRequest) *UpdateEquipmentEsnInvoker {
	requestDef := GenReqDefForUpdateEquipmentEsn()
	return &UpdateEquipmentEsnInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEquipmentInfo 更新智能企业网关设备
//
// 更新智能企业网关设备
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) UpdateEquipmentInfo(request *model.UpdateEquipmentInfoRequest) (*model.UpdateEquipmentInfoResponse, error) {
	requestDef := GenReqDefForUpdateEquipmentInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEquipmentInfoResponse), nil
	}
}

// UpdateEquipmentInfoInvoker 更新智能企业网关设备
func (c *EcClient) UpdateEquipmentInfoInvoker(request *model.UpdateEquipmentInfoRequest) *UpdateEquipmentInfoInvoker {
	requestDef := GenReqDefForUpdateEquipmentInfo()
	return &UpdateEquipmentInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEquipmentLanConfig 创建智能企业网关设备LAN口配置
//
// 创建智能企业网关设备LAN口配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) CreateEquipmentLanConfig(request *model.CreateEquipmentLanConfigRequest) (*model.CreateEquipmentLanConfigResponse, error) {
	requestDef := GenReqDefForCreateEquipmentLanConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEquipmentLanConfigResponse), nil
	}
}

// CreateEquipmentLanConfigInvoker 创建智能企业网关设备LAN口配置
func (c *EcClient) CreateEquipmentLanConfigInvoker(request *model.CreateEquipmentLanConfigRequest) *CreateEquipmentLanConfigInvoker {
	requestDef := GenReqDefForCreateEquipmentLanConfig()
	return &CreateEquipmentLanConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEquipmentLanConfig 删除智能企业网关设备LAN口配置
//
// 删除智能企业网关设备LAN口配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) DeleteEquipmentLanConfig(request *model.DeleteEquipmentLanConfigRequest) (*model.DeleteEquipmentLanConfigResponse, error) {
	requestDef := GenReqDefForDeleteEquipmentLanConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEquipmentLanConfigResponse), nil
	}
}

// DeleteEquipmentLanConfigInvoker 删除智能企业网关设备LAN口配置
func (c *EcClient) DeleteEquipmentLanConfigInvoker(request *model.DeleteEquipmentLanConfigRequest) *DeleteEquipmentLanConfigInvoker {
	requestDef := GenReqDefForDeleteEquipmentLanConfig()
	return &DeleteEquipmentLanConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEquipmentInterfaceName 查询智能企业网关已配置的接口名字
//
// 查询智能企业网关已配置的接口名字
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ListEquipmentInterfaceName(request *model.ListEquipmentInterfaceNameRequest) (*model.ListEquipmentInterfaceNameResponse, error) {
	requestDef := GenReqDefForListEquipmentInterfaceName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEquipmentInterfaceNameResponse), nil
	}
}

// ListEquipmentInterfaceNameInvoker 查询智能企业网关已配置的接口名字
func (c *EcClient) ListEquipmentInterfaceNameInvoker(request *model.ListEquipmentInterfaceNameRequest) *ListEquipmentInterfaceNameInvoker {
	requestDef := GenReqDefForListEquipmentInterfaceName()
	return &ListEquipmentInterfaceNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEquipmentDnsInfo 查询智能企业网关设备主备DNS配置
//
// 查询智能企业网关设备主备DNS配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ShowEquipmentDnsInfo(request *model.ShowEquipmentDnsInfoRequest) (*model.ShowEquipmentDnsInfoResponse, error) {
	requestDef := GenReqDefForShowEquipmentDnsInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEquipmentDnsInfoResponse), nil
	}
}

// ShowEquipmentDnsInfoInvoker 查询智能企业网关设备主备DNS配置
func (c *EcClient) ShowEquipmentDnsInfoInvoker(request *model.ShowEquipmentDnsInfoRequest) *ShowEquipmentDnsInfoInvoker {
	requestDef := GenReqDefForShowEquipmentDnsInfo()
	return &ShowEquipmentDnsInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEquipmentLanInfo 查询智能企业网关设备LAN口配置
//
// 查询智能企业网关设备LAN口配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ShowEquipmentLanInfo(request *model.ShowEquipmentLanInfoRequest) (*model.ShowEquipmentLanInfoResponse, error) {
	requestDef := GenReqDefForShowEquipmentLanInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEquipmentLanInfoResponse), nil
	}
}

// ShowEquipmentLanInfoInvoker 查询智能企业网关设备LAN口配置
func (c *EcClient) ShowEquipmentLanInfoInvoker(request *model.ShowEquipmentLanInfoRequest) *ShowEquipmentLanInfoInvoker {
	requestDef := GenReqDefForShowEquipmentLanInfo()
	return &ShowEquipmentLanInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEquipmentDnsInfo 更新智能企业网关设备主备DNS配置
//
// 更新智能企业网关设备主备DNS配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) UpdateEquipmentDnsInfo(request *model.UpdateEquipmentDnsInfoRequest) (*model.UpdateEquipmentDnsInfoResponse, error) {
	requestDef := GenReqDefForUpdateEquipmentDnsInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEquipmentDnsInfoResponse), nil
	}
}

// UpdateEquipmentDnsInfoInvoker 更新智能企业网关设备主备DNS配置
func (c *EcClient) UpdateEquipmentDnsInfoInvoker(request *model.UpdateEquipmentDnsInfoRequest) *UpdateEquipmentDnsInfoInvoker {
	requestDef := GenReqDefForUpdateEquipmentDnsInfo()
	return &UpdateEquipmentDnsInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEquipmentLanConfig 更新智能企业网关设备LAN口配置
//
// 更新智能企业网关设备LAN口配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) UpdateEquipmentLanConfig(request *model.UpdateEquipmentLanConfigRequest) (*model.UpdateEquipmentLanConfigResponse, error) {
	requestDef := GenReqDefForUpdateEquipmentLanConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEquipmentLanConfigResponse), nil
	}
}

// UpdateEquipmentLanConfigInvoker 更新智能企业网关设备LAN口配置
func (c *EcClient) UpdateEquipmentLanConfigInvoker(request *model.UpdateEquipmentLanConfigRequest) *UpdateEquipmentLanConfigInvoker {
	requestDef := GenReqDefForUpdateEquipmentLanConfig()
	return &UpdateEquipmentLanConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEquipmentOspf 查询智能企业网关设备OSPF配置
//
// 查询智能企业网关设备OSPF配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ShowEquipmentOspf(request *model.ShowEquipmentOspfRequest) (*model.ShowEquipmentOspfResponse, error) {
	requestDef := GenReqDefForShowEquipmentOspf()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEquipmentOspfResponse), nil
	}
}

// ShowEquipmentOspfInvoker 查询智能企业网关设备OSPF配置
func (c *EcClient) ShowEquipmentOspfInvoker(request *model.ShowEquipmentOspfRequest) *ShowEquipmentOspfInvoker {
	requestDef := GenReqDefForShowEquipmentOspf()
	return &ShowEquipmentOspfInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEquipmentOspf 配置智能企业网关设备OSPF协议
//
// 配置智能企业网关设备OSPF协议
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) UpdateEquipmentOspf(request *model.UpdateEquipmentOspfRequest) (*model.UpdateEquipmentOspfResponse, error) {
	requestDef := GenReqDefForUpdateEquipmentOspf()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEquipmentOspfResponse), nil
	}
}

// UpdateEquipmentOspfInvoker 配置智能企业网关设备OSPF协议
func (c *EcClient) UpdateEquipmentOspfInvoker(request *model.UpdateEquipmentOspfRequest) *UpdateEquipmentOspfInvoker {
	requestDef := GenReqDefForUpdateEquipmentOspf()
	return &UpdateEquipmentOspfInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEquipmentStaticRouteConfig 创建智能企业网关设备静态路由配置
//
// 创建智能企业网关设备静态路由配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) CreateEquipmentStaticRouteConfig(request *model.CreateEquipmentStaticRouteConfigRequest) (*model.CreateEquipmentStaticRouteConfigResponse, error) {
	requestDef := GenReqDefForCreateEquipmentStaticRouteConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEquipmentStaticRouteConfigResponse), nil
	}
}

// CreateEquipmentStaticRouteConfigInvoker 创建智能企业网关设备静态路由配置
func (c *EcClient) CreateEquipmentStaticRouteConfigInvoker(request *model.CreateEquipmentStaticRouteConfigRequest) *CreateEquipmentStaticRouteConfigInvoker {
	requestDef := GenReqDefForCreateEquipmentStaticRouteConfig()
	return &CreateEquipmentStaticRouteConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEquipmentStaticRouteConfig 删除智能企业网关设备静态路由配置
//
// 删除智能企业网关设备静态路由配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) DeleteEquipmentStaticRouteConfig(request *model.DeleteEquipmentStaticRouteConfigRequest) (*model.DeleteEquipmentStaticRouteConfigResponse, error) {
	requestDef := GenReqDefForDeleteEquipmentStaticRouteConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEquipmentStaticRouteConfigResponse), nil
	}
}

// DeleteEquipmentStaticRouteConfigInvoker 删除智能企业网关设备静态路由配置
func (c *EcClient) DeleteEquipmentStaticRouteConfigInvoker(request *model.DeleteEquipmentStaticRouteConfigRequest) *DeleteEquipmentStaticRouteConfigInvoker {
	requestDef := GenReqDefForDeleteEquipmentStaticRouteConfig()
	return &DeleteEquipmentStaticRouteConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEquipmentStaticRouteInfo 查询智能企业网关设备静态路由配置
//
// 查询智能企业网关设备静态路由配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ShowEquipmentStaticRouteInfo(request *model.ShowEquipmentStaticRouteInfoRequest) (*model.ShowEquipmentStaticRouteInfoResponse, error) {
	requestDef := GenReqDefForShowEquipmentStaticRouteInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEquipmentStaticRouteInfoResponse), nil
	}
}

// ShowEquipmentStaticRouteInfoInvoker 查询智能企业网关设备静态路由配置
func (c *EcClient) ShowEquipmentStaticRouteInfoInvoker(request *model.ShowEquipmentStaticRouteInfoRequest) *ShowEquipmentStaticRouteInfoInvoker {
	requestDef := GenReqDefForShowEquipmentStaticRouteInfo()
	return &ShowEquipmentStaticRouteInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEquipmentStaticRouteConfig 更新智能企业网关设备静态路由配置
//
// 更新智能企业网关设备静态路由配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) UpdateEquipmentStaticRouteConfig(request *model.UpdateEquipmentStaticRouteConfigRequest) (*model.UpdateEquipmentStaticRouteConfigResponse, error) {
	requestDef := GenReqDefForUpdateEquipmentStaticRouteConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEquipmentStaticRouteConfigResponse), nil
	}
}

// UpdateEquipmentStaticRouteConfigInvoker 更新智能企业网关设备静态路由配置
func (c *EcClient) UpdateEquipmentStaticRouteConfigInvoker(request *model.UpdateEquipmentStaticRouteConfigRequest) *UpdateEquipmentStaticRouteConfigInvoker {
	requestDef := GenReqDefForUpdateEquipmentStaticRouteConfig()
	return &UpdateEquipmentStaticRouteConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEquipmentWanInfo 查询智能企业网关设备WAN口配置
//
// 查询智能企业网关设备WAN口配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ShowEquipmentWanInfo(request *model.ShowEquipmentWanInfoRequest) (*model.ShowEquipmentWanInfoResponse, error) {
	requestDef := GenReqDefForShowEquipmentWanInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEquipmentWanInfoResponse), nil
	}
}

// ShowEquipmentWanInfoInvoker 查询智能企业网关设备WAN口配置
func (c *EcClient) ShowEquipmentWanInfoInvoker(request *model.ShowEquipmentWanInfoRequest) *ShowEquipmentWanInfoInvoker {
	requestDef := GenReqDefForShowEquipmentWanInfo()
	return &ShowEquipmentWanInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEquipmentWanConfig 更新智能企业网关设备WAN口配置
//
// 更新智能企业网关设备WAN口配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) UpdateEquipmentWanConfig(request *model.UpdateEquipmentWanConfigRequest) (*model.UpdateEquipmentWanConfigResponse, error) {
	requestDef := GenReqDefForUpdateEquipmentWanConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEquipmentWanConfigResponse), nil
	}
}

// UpdateEquipmentWanConfigInvoker 更新智能企业网关设备WAN口配置
func (c *EcClient) UpdateEquipmentWanConfigInvoker(request *model.UpdateEquipmentWanConfigRequest) *UpdateEquipmentWanConfigInvoker {
	requestDef := GenReqDefForUpdateEquipmentWanConfig()
	return &UpdateEquipmentWanConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddEcnWithEr 关联企业路由器到企业连接网络
//
// 关联企业路由器到企业连接网络
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) AddEcnWithEr(request *model.AddEcnWithErRequest) (*model.AddEcnWithErResponse, error) {
	requestDef := GenReqDefForAddEcnWithEr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddEcnWithErResponse), nil
	}
}

// AddEcnWithErInvoker 关联企业路由器到企业连接网络
func (c *EcClient) AddEcnWithErInvoker(request *model.AddEcnWithErRequest) *AddEcnWithErInvoker {
	requestDef := GenReqDefForAddEcnWithEr()
	return &AddEcnWithErInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEcnWithEr 解除企业路由器和企业连接网络的关联
//
// 解除企业路由器和企业连接网络的关联
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) DeleteEcnWithEr(request *model.DeleteEcnWithErRequest) (*model.DeleteEcnWithErResponse, error) {
	requestDef := GenReqDefForDeleteEcnWithEr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEcnWithErResponse), nil
	}
}

// DeleteEcnWithErInvoker 解除企业路由器和企业连接网络的关联
func (c *EcClient) DeleteEcnWithErInvoker(request *model.DeleteEcnWithErRequest) *DeleteEcnWithErInvoker {
	requestDef := GenReqDefForDeleteEcnWithEr()
	return &DeleteEcnWithErInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEcnWithEr 查询企业连接网络网与企业路由器关联关系
//
// 根据企业连接网络ID，查询企业连接网络网与企业路由器关联关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ListEcnWithEr(request *model.ListEcnWithErRequest) (*model.ListEcnWithErResponse, error) {
	requestDef := GenReqDefForListEcnWithEr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEcnWithErResponse), nil
	}
}

// ListEcnWithErInvoker 查询企业连接网络网与企业路由器关联关系
func (c *EcClient) ListEcnWithErInvoker(request *model.ListEcnWithErRequest) *ListEcnWithErInvoker {
	requestDef := GenReqDefForListEcnWithEr()
	return &ListEcnWithErInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeIegPassword 修改IEG设备admin账户密码
//
// 修改IEG设备admin账户密码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ChangeIegPassword(request *model.ChangeIegPasswordRequest) (*model.ChangeIegPasswordResponse, error) {
	requestDef := GenReqDefForChangeIegPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeIegPasswordResponse), nil
	}
}

// ChangeIegPasswordInvoker 修改IEG设备admin账户密码
func (c *EcClient) ChangeIegPasswordInvoker(request *model.ChangeIegPasswordRequest) *ChangeIegPasswordInvoker {
	requestDef := GenReqDefForChangeIegPassword()
	return &ChangeIegPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIeg 查询租户智能企业网关列表
//
// 查询租户智能企业网关列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ListIeg(request *model.ListIegRequest) (*model.ListIegResponse, error) {
	requestDef := GenReqDefForListIeg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIegResponse), nil
	}
}

// ListIegInvoker 查询租户智能企业网关列表
func (c *EcClient) ListIegInvoker(request *model.ListIegRequest) *ListIegInvoker {
	requestDef := GenReqDefForListIeg()
	return &ListIegInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIegInfo 查询租户单个智能企业网关
//
// 根据智能企业网关ID，查询租户智能企业网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ShowIegInfo(request *model.ShowIegInfoRequest) (*model.ShowIegInfoResponse, error) {
	requestDef := GenReqDefForShowIegInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIegInfoResponse), nil
	}
}

// ShowIegInfoInvoker 查询租户单个智能企业网关
func (c *EcClient) ShowIegInfoInvoker(request *model.ShowIegInfoRequest) *ShowIegInfoInvoker {
	requestDef := GenReqDefForShowIegInfo()
	return &ShowIegInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchEquipmentHaType 交换双机主备属性
//
// 交换双机主备属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) SwitchEquipmentHaType(request *model.SwitchEquipmentHaTypeRequest) (*model.SwitchEquipmentHaTypeResponse, error) {
	requestDef := GenReqDefForSwitchEquipmentHaType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchEquipmentHaTypeResponse), nil
	}
}

// SwitchEquipmentHaTypeInvoker 交换双机主备属性
func (c *EcClient) SwitchEquipmentHaTypeInvoker(request *model.SwitchEquipmentHaTypeRequest) *SwitchEquipmentHaTypeInvoker {
	requestDef := GenReqDefForSwitchEquipmentHaType()
	return &SwitchEquipmentHaTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIeg 更新智能企业网关
//
// 根据智能企业网关ID，更新智能企业网关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) UpdateIeg(request *model.UpdateIegRequest) (*model.UpdateIegResponse, error) {
	requestDef := GenReqDefForUpdateIeg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIegResponse), nil
	}
}

// UpdateIegInvoker 更新智能企业网关
func (c *EcClient) UpdateIegInvoker(request *model.UpdateIegRequest) *UpdateIegInvoker {
	requestDef := GenReqDefForUpdateIeg()
	return &UpdateIegInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuotasInfo 查询EC相关的指定租户的配额
//
// 查询EC相关的指定租户的配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ShowQuotasInfo(request *model.ShowQuotasInfoRequest) (*model.ShowQuotasInfoResponse, error) {
	requestDef := GenReqDefForShowQuotasInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasInfoResponse), nil
	}
}

// ShowQuotasInfoInvoker 查询EC相关的指定租户的配额
func (c *EcClient) ShowQuotasInfoInvoker(request *model.ShowQuotasInfoRequest) *ShowQuotasInfoInvoker {
	requestDef := GenReqDefForShowQuotasInfo()
	return &ShowQuotasInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddVrrpConfig 创建vrrp配置
//
// 创建vrrp配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) AddVrrpConfig(request *model.AddVrrpConfigRequest) (*model.AddVrrpConfigResponse, error) {
	requestDef := GenReqDefForAddVrrpConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddVrrpConfigResponse), nil
	}
}

// AddVrrpConfigInvoker 创建vrrp配置
func (c *EcClient) AddVrrpConfigInvoker(request *model.AddVrrpConfigRequest) *AddVrrpConfigInvoker {
	requestDef := GenReqDefForAddVrrpConfig()
	return &AddVrrpConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVrrpConfig 删除vrrp配置
//
// 删除vrrp配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) DeleteVrrpConfig(request *model.DeleteVrrpConfigRequest) (*model.DeleteVrrpConfigResponse, error) {
	requestDef := GenReqDefForDeleteVrrpConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVrrpConfigResponse), nil
	}
}

// DeleteVrrpConfigInvoker 删除vrrp配置
func (c *EcClient) DeleteVrrpConfigInvoker(request *model.DeleteVrrpConfigRequest) *DeleteVrrpConfigInvoker {
	requestDef := GenReqDefForDeleteVrrpConfig()
	return &DeleteVrrpConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVrrpConfig 查询vrrp配置列表
//
// 查询vrrp配置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) ShowVrrpConfig(request *model.ShowVrrpConfigRequest) (*model.ShowVrrpConfigResponse, error) {
	requestDef := GenReqDefForShowVrrpConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVrrpConfigResponse), nil
	}
}

// ShowVrrpConfigInvoker 查询vrrp配置列表
func (c *EcClient) ShowVrrpConfigInvoker(request *model.ShowVrrpConfigRequest) *ShowVrrpConfigInvoker {
	requestDef := GenReqDefForShowVrrpConfig()
	return &ShowVrrpConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVrrpConfig 更新vrrp配置
//
// 更新vrrp配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EcClient) UpdateVrrpConfig(request *model.UpdateVrrpConfigRequest) (*model.UpdateVrrpConfigResponse, error) {
	requestDef := GenReqDefForUpdateVrrpConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVrrpConfigResponse), nil
	}
}

// UpdateVrrpConfigInvoker 更新vrrp配置
func (c *EcClient) UpdateVrrpConfigInvoker(request *model.UpdateVrrpConfigRequest) *UpdateVrrpConfigInvoker {
	requestDef := GenReqDefForUpdateVrrpConfig()
	return &UpdateVrrpConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
