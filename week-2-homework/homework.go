package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

/*
问题：
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。
为什么，应该怎么做请写出代码？

回答：
ErrNoRows产生于当QueryRow方法不返回行时，推迟到Scan方法返回ErrNoRows。
我们在遇到这个错误时，应当调用Wrap方法，捕获错误的堆栈信息，即时定位错误，告知上层需要进行修改。

*/

func main() {
	err := dao()
	fmt.Printf("dao err:%+v", err)
}

func dao() error {
	return errors.Wrap(sql.ErrNoRows, "dao error")
}
