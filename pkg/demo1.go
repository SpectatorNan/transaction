package pkg

import "fmt"

func Demo1() (err error) {
	beginTransaction()
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
			fmt.Printf("%v panic\n", e)
			rollback()
		}
	}()
	if err = one(); err != nil {
		rollback()
		return err
	}
	if err = two(); err != nil {
		rollback()
		return err
	}
	if err = three(); err != nil {
		rollback()
		return err
	}
	if err = four(); err != nil {
		rollback()
		return err
	}
	if err = five(); err != nil {
		rollback()
		return err
	}
	commit()
	return nil
}
