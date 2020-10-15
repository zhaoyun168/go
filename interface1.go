package main

import "fmt"

type PayMethod interface {
	pay()
}

type wechat struct {
}

func (wechatPay wechat) pay() {
	fmt.Println("wechat pay")
}

type al struct{
}

func (alpay al) pay() {
	fmt.Println("al pay")
}

func main() {
	var pay_method PayMethod

	pay_method = new(wechat)
	pay_method.pay()

	pay_method = new(al)
	pay_method.pay()
}