package pkg

import (
	"errors"
	"fmt"
)

//开启事务
func beginTransaction() {
	fmt.Println("beginTransaction")
}

//回滚事务
func rollback() {
	fmt.Println("rollback")
}

//提交事务
func commit() {
	fmt.Println("commit")
}

//执行one操作
func one() (err error) {
	fmt.Println("one ok")
	return nil
}

//执行two操作
func two() (err error) {
	fmt.Println("two ok")
	return nil
}

//执行three操作
func three() (err error) {
	fmt.Println("two ok")
	return nil
}

//执行four操作
func four() (err error) {
	fmt.Println("four ok")
	return nil
}

//执行five操作
func five() (err error) {
	err = errors.New("five panic")
	panic("five")
	return err
}
