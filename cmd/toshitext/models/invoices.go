package models

type Invoice struct {
	ID            int
	Currency      int
	PaymentAmount int
}

// func alsonotmain() {

// }

// func (i *Invoice) Create() error {
// 	// creates invoice in db
// 	err := db.create()
// 	if err != nil {
// 		return err
// 	}
// }
