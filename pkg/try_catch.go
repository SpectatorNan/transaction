package pkg

import (
	"fmt"
	"transaction/exception"
)

func Demo4() (err error) {
	exception.Block{
		Try: func() {
			beginTransaction()
			if err = one(); err != nil {
				panic(err)
			}
			if err = two(); err != nil {
				panic(err)
			}
			if err = three(); err != nil {
				panic(err)
			}
			if err = four(); err != nil {
				panic(err)
			}
			if err = five(); err != nil {
				panic(err)
			}
			err = nil
			commit()
		},
		Catch: func(e interface{}) {
			rollback()
			fmt.Printf("%v panic\n", e)
			err = fmt.Errorf("%v", e)
		},
	}.Do()
	return err
}
