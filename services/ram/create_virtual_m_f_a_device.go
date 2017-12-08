
package ram

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

func (client *Client) CreateVirtualMFADevice(request *CreateVirtualMFADeviceRequest) (response *CreateVirtualMFADeviceResponse, err error) {
response = CreateCreateVirtualMFADeviceResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) CreateVirtualMFADeviceWithChan(request *CreateVirtualMFADeviceRequest) (<-chan *CreateVirtualMFADeviceResponse, <-chan error) {
responseChan := make(chan *CreateVirtualMFADeviceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CreateVirtualMFADevice(request)
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

func (client *Client) CreateVirtualMFADeviceWithCallback(request *CreateVirtualMFADeviceRequest, callback func(response *CreateVirtualMFADeviceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CreateVirtualMFADeviceResponse
var err error
defer close(result)
response, err = client.CreateVirtualMFADevice(request)
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

type CreateVirtualMFADeviceRequest struct {
*requests.RpcRequest
                VirtualMFADeviceName  string `position:"Query" name:"VirtualMFADeviceName"`
}


type CreateVirtualMFADeviceResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            VirtualMFADevice struct {
            SerialNumber     string `json:"SerialNumber" xml:"SerialNumber"`
            Base32StringSeed     string `json:"Base32StringSeed" xml:"Base32StringSeed"`
            QRCodePNG     string `json:"QRCodePNG" xml:"QRCodePNG"`
            }  `json:"VirtualMFADevice" xml:"VirtualMFADevice"`
}

func CreateCreateVirtualMFADeviceRequest() (request *CreateVirtualMFADeviceRequest) {
request = &CreateVirtualMFADeviceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ram", "2015-05-01", "CreateVirtualMFADevice", "", "")
return
}

func CreateCreateVirtualMFADeviceResponse() (response *CreateVirtualMFADeviceResponse) {
response = &CreateVirtualMFADeviceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}
