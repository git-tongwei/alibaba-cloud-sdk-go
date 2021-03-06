package dm

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

func (client *Client) DeleteSign(request *DeleteSignRequest) (response *DeleteSignResponse, err error) {
	response = CreateDeleteSignResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteSignWithChan(request *DeleteSignRequest) (<-chan *DeleteSignResponse, <-chan error) {
	responseChan := make(chan *DeleteSignResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSign(request)
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

func (client *Client) DeleteSignWithCallback(request *DeleteSignRequest, callback func(response *DeleteSignResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSignResponse
		var err error
		defer close(result)
		response, err = client.DeleteSign(request)
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

type DeleteSignRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	SignId               string `position:"Query" name:"SignId"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	FromType             string `position:"Query" name:"FromType"`
}

type DeleteSignResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteSignRequest() (request *DeleteSignRequest) {
	request = &DeleteSignRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "DeleteSign", "", "")
	return
}

func CreateDeleteSignResponse() (response *DeleteSignResponse) {
	response = &DeleteSignResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
