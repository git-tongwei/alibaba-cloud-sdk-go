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

func (client *Client) DeleteServerCertificate(request *DeleteServerCertificateRequest) (response *DeleteServerCertificateResponse, err error) {
	response = CreateDeleteServerCertificateResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteServerCertificateWithChan(request *DeleteServerCertificateRequest) (<-chan *DeleteServerCertificateResponse, <-chan error) {
	responseChan := make(chan *DeleteServerCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteServerCertificate(request)
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

func (client *Client) DeleteServerCertificateWithCallback(request *DeleteServerCertificateRequest, callback func(response *DeleteServerCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteServerCertificateResponse
		var err error
		defer close(result)
		response, err = client.DeleteServerCertificate(request)
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

type DeleteServerCertificateRequest struct {
	*requests.RpcRequest
	Tags                 string `position:"Query" name:"Tags"`
	ServerCertificateId  string `position:"Query" name:"ServerCertificateId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AccessKeyId          string `position:"Query" name:"access_key_id"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type DeleteServerCertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteServerCertificateRequest() (request *DeleteServerCertificateRequest) {
	request = &DeleteServerCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DeleteServerCertificate", "", "")
	return
}

func CreateDeleteServerCertificateResponse() (response *DeleteServerCertificateResponse) {
	response = &DeleteServerCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
