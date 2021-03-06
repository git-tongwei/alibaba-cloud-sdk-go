package slb

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) SetLoadBalancerUDPListenerAttribute(request *SetLoadBalancerUDPListenerAttributeRequest) (response *SetLoadBalancerUDPListenerAttributeResponse, err error) {
	response = CreateSetLoadBalancerUDPListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SetLoadBalancerUDPListenerAttributeWithChan(request *SetLoadBalancerUDPListenerAttributeRequest) (<-chan *SetLoadBalancerUDPListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *SetLoadBalancerUDPListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetLoadBalancerUDPListenerAttribute(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) SetLoadBalancerUDPListenerAttributeWithCallback(request *SetLoadBalancerUDPListenerAttributeRequest, callback func(response *SetLoadBalancerUDPListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetLoadBalancerUDPListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.SetLoadBalancerUDPListenerAttribute(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type SetLoadBalancerUDPListenerAttributeRequest struct {
	*requests.RpcRequest
	VServerGroup              string `position:"Query" name:"VServerGroup"`
	UnhealthyThreshold        string `position:"Query" name:"UnhealthyThreshold"`
	Bandwidth                 string `position:"Query" name:"Bandwidth"`
	ResourceOwnerAccount      string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId           string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount              string `position:"Query" name:"OwnerAccount"`
	MasterSlaveServerGroupId  string `position:"Query" name:"MasterSlaveServerGroupId"`
	PersistenceTimeout        string `position:"Query" name:"PersistenceTimeout"`
	Tags                      string `position:"Query" name:"Tags"`
	Scheduler                 string `position:"Query" name:"Scheduler"`
	OwnerId                   string `position:"Query" name:"OwnerId"`
	VServerGroupId            string `position:"Query" name:"VServerGroupId"`
	HealthCheckInterval       string `position:"Query" name:"HealthCheckInterval"`
	ListenerPort              string `position:"Query" name:"ListenerPort"`
	HealthCheckExp            string `position:"Query" name:"healthCheckExp"`
	MaxConnection             string `position:"Query" name:"MaxConnection"`
	AccessKeyId               string `position:"Query" name:"access_key_id"`
	HealthCheckReq            string `position:"Query" name:"healthCheckReq"`
	HealthCheckConnectPort    string `position:"Query" name:"HealthCheckConnectPort"`
	LoadBalancerId            string `position:"Query" name:"LoadBalancerId"`
	MasterSlaveServerGroup    string `position:"Query" name:"MasterSlaveServerGroup"`
	HealthyThreshold          string `position:"Query" name:"HealthyThreshold"`
	HealthCheckConnectTimeout string `position:"Query" name:"HealthCheckConnectTimeout"`
}

type SetLoadBalancerUDPListenerAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateSetLoadBalancerUDPListenerAttributeRequest() (request *SetLoadBalancerUDPListenerAttributeRequest) {
	request = &SetLoadBalancerUDPListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "SetLoadBalancerUDPListenerAttribute", "", "")
	return
}

func CreateSetLoadBalancerUDPListenerAttributeResponse() (response *SetLoadBalancerUDPListenerAttributeResponse) {
	response = &SetLoadBalancerUDPListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
