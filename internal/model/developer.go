package model

import (
	"fmt"
	"strconv"
)

type Developer struct {
	AbstractEmployee
	Overtime uint8
}

func (d Developer) Info() AbstractEmployee {
	return d.AbstractEmployee
}

func (d Developer) Salary() float64 {
	return float64(d.BasePay) + (float64(d.Overtime) * 200000)
}

func (d Developer) String() string {
	return fmt.Sprint("{" +
		"ID:" + strconv.FormatInt(d.ID, 10) + ", " +
		"Name:" + d.Fullname + ", " +
		"Age:" + fmt.Sprintf("%d", d.Age) + ", " +
		"PhoneNo:" + d.PhoneNo + ", " +
		"Email:" + d.Email + ", " +
		"BasePay:" + fmt.Sprintf("%0.00f", d.BasePay) + ", " +
		"Overtime:" + fmt.Sprintf("%d", d.Overtime) + ", " +
		"}",
	)
}
