package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/model"
)

type VpcClient struct {
	hcClient *http_client.HcHttpClient
}

func NewVpcClient(hcClient *http_client.HcHttpClient) *VpcClient {
	return &VpcClient{hcClient: hcClient}
}

func VpcClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//批量创建辅助弹性网卡
func (c *VpcClient) BatchCreateSubNetworkInterfaceV3(request *model.BatchCreateSubNetworkInterfaceV3Request) (*model.BatchCreateSubNetworkInterfaceV3Response, error) {
	requestDef := GenReqDefForBatchCreateSubNetworkInterfaceV3(request)
	resp, responseDef := GenRespForBatchCreateSubNetworkInterfaceV3()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建辅助弹性网卡
func (c *VpcClient) CreateSubNetworkInterface(request *model.CreateSubNetworkInterfaceRequest) (*model.CreateSubNetworkInterfaceResponse, error) {
	requestDef := GenReqDefForCreateSubNetworkInterface(request)
	resp, responseDef := GenRespForCreateSubNetworkInterface()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除辅助弹性网卡
func (c *VpcClient) DeleteSubNetworkInterface(request *model.DeleteSubNetworkInterfaceRequest) (*model.DeleteSubNetworkInterfaceResponse, error) {
	requestDef := GenReqDefForDeleteSubNetworkInterface(request)
	resp, responseDef := GenRespForDeleteSubNetworkInterface()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询辅助弹性网卡列表，单次查询最多返回2000条数据
func (c *VpcClient) ListSubNetworkInterfaces(request *model.ListSubNetworkInterfacesRequest) (*model.ListSubNetworkInterfacesResponse, error) {
	requestDef := GenReqDefForListSubNetworkInterfaces(request)
	resp, responseDef := GenRespForListSubNetworkInterfaces()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询辅助弹性网卡详情
func (c *VpcClient) ShowSubNetworkInterface(request *model.ShowSubNetworkInterfaceRequest) (*model.ShowSubNetworkInterfaceResponse, error) {
	requestDef := GenReqDefForShowSubNetworkInterface(request)
	resp, responseDef := GenRespForShowSubNetworkInterface()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询辅助弹性网卡数目
func (c *VpcClient) ShowSubNetworkInterfacesQuantity(request *model.ShowSubNetworkInterfacesQuantityRequest) (*model.ShowSubNetworkInterfacesQuantityResponse, error) {
	requestDef := GenReqDefForShowSubNetworkInterfacesQuantity(request)
	resp, responseDef := GenRespForShowSubNetworkInterfacesQuantity()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新辅助弹性网卡
func (c *VpcClient) UpdateSubNetworkInterface(request *model.UpdateSubNetworkInterfaceRequest) (*model.UpdateSubNetworkInterfaceResponse, error) {
	requestDef := GenReqDefForUpdateSubNetworkInterface(request)
	resp, responseDef := GenRespForUpdateSubNetworkInterface()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
