package integration

import (
	"fmt"
	"reflect"
)

// 定义接口
type MyInterface interface {
	MyMethod() string
}

type TContract interface {
}

type Client[T TContract] struct { // 含有conn,grpc.call方法
}

func NewMyInterface[T TContract]() MyInterface {
	interfaceType := reflect.TypeOf((*MyInterface)(nil)).Elem()

	// 创建结构体类型
	structType := reflect.StructOf([]reflect.StructField{
		{
			Name: "integration.myField",
			Type: reflect.TypeOf((T)(nil)),
		},
	})

	// 定义实现接口中的方法
	myMethodFunc := func(args []reflect.Value) []reflect.Value {
		return []reflect.Value{reflect.ValueOf("Hello, world!")}
	}

	// 使用反射来构建实现该接口的结构体类型
	method := reflect.Method{
		Name: "MyMethod",
		Func: reflect.MakeFunc(interfaceType.Method(0).Type, myMethodFunc),
	}
	structType = reflect.StructOf([]reflect.StructField{
		{
			Name: "myField",
			Type: reflect.TypeOf(""),
		},
		{
			Name: "MyMethod",
			Type: method.Type,
		},
	})
	implType := reflect.PtrTo(structType)

	fmt.Println("type:===", implType)

	// 创建实现该接口的结构体类型的实例
	implValue := reflect.New(structType).Elem()
	implValue.Field(1).Set(method.Func)

	// 将实现该接口的结构体类型的实例转换为接口类型的值
	i := implValue.Addr().Interface().(MyInterface)

	// 调用实现的方法
	fmt.Println(i.MyMethod())
	return i
}
