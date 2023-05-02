package model

import (
	"fmt"
	"strconv"
)

type Tester struct {
	AbstractEmployee
	BugsFount int64
}

func (t Tester) Info() AbstractEmployee {
	return t.AbstractEmployee
}

func (t Tester) Salary() float64 {
	return float64(t.BasePay) + (float64(t.BugsFount) * 50000)
}

func (t Tester) String() string {
	return fmt.Sprint("{" +
		"It:" + strconv.FormatInt(t.ID, 10) + ", " +
		"Name:" + t.Fullname + ", " +
		"Age:" + fmt.Sprintf("%d", t.Age) + ", " +
		"PhoneNo:" + t.PhoneNo + ", " +
		"Email:" + t.Email + ", " +
		"BasePay:" + fmt.Sprintf("%0.00f", t.BasePay) + ", " +
		"BugsFound:" + string(rune(t.BugsFount)) + ", " +
		"}",
	)
}
