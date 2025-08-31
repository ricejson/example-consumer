package impl

import (
	"github.com/ricejson/example-common/model"
	model2 "github.com/ricejson/rice-rpc-easy/model"
	"github.com/ricejson/rice-rpc-easy/serializer"
	"io"
	"log"
	"net/http"
	"strings"
)

// UserServiceImpl 静态代理实现
type UserServiceImpl struct {
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{}
}

func (u *UserServiceImpl) GetUser() (model.User, error) {
	// 创建序列化器
	nativeSerializer := serializer.NewNativeSerializer()
	// 发送HTTP请求
	request := model2.RpcRequest{}
	request.ServiceName = "UserService"
	request.MethodName = "GetUser"
	// 序列化
	bytes, _ := nativeSerializer.Serialize(request)
	response, _ := http.Post("http://localhost:8080/", "application/json", strings.NewReader(string(bytes)))
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	// 反序列化
	rpcResponse := model2.RpcResponse{}
	_ = nativeSerializer.Deserialize(body, &rpcResponse)
	log.Println(rpcResponse)
	var user model.User
	serialize, _ := nativeSerializer.Serialize(rpcResponse.Data)
	nativeSerializer.Deserialize(serialize, &user)
	return user, nil
}
