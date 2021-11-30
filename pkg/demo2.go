package pkg

import "fmt"

func Demo2() (err error) {
	beginTransaction()
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
			fmt.Printf("%v panic\n", e)
			rollback()
		}
	}()
	if err = one(); err != nil {
		goto ROLLBACK
	}
	if err = two(); err != nil {
		goto ROLLBACK
	}
	if err = three(); err != nil {
		goto ROLLBACK
	}
	if err = four(); err != nil {
		goto ROLLBACK
	}
	if err = five(); err != nil {
		goto ROLLBACK
	}
	commit()
	return nil
ROLLBACK:
	rollback()
	return err
}
