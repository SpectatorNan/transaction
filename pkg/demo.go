package pkg

import "fmt"

func Demo() (err error) {
	beginTransaction()
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
			fmt.Printf("%v panic\n", e)
			rollback()
		}
	}()
	if err = one(); err == nil {
		if err = two(); err == nil {
			if err = three(); err == nil {
				if err = four(); err == nil {
					if err = five(); err == nil {
						commit()
						return nil
					} else {
						rollback()
						return err
					}
				} else {
					rollback()
					return err
				}
			} else {
				rollback()
				return err
			}
		} else {
			rollback()
			return err
		}
	} else {
		rollback()
		return err
	}
}
