package main

import (
	"fmt"

	"virsavik.github.io/employee-manager/internal/model"
	"virsavik.github.io/employee-manager/pkg/linkedlist"
	"virsavik.github.io/employee-manager/pkg/ops"
)

func main() {
	employeeList := linkedlist.OfArray([]model.Employee{
		model.Developer{
			AbstractEmployee: model.AbstractEmployee{
				ID:       1,
				Fullname: "Virsavik",
				Age:      22,
				PhoneNo:  "0987654321",
				Email:    "vir@sample.com",
				BasePay:  100.0,
			},
			Overtime: 1,
		},
		model.Tester{
			AbstractEmployee: model.AbstractEmployee{
				ID:       2,
				Fullname: "Yenefer",
				Age:      22,
				PhoneNo:  "0987654322",
				Email:    "yen@sample.com",
				BasePay:  900.0,
			},
			BugsFount: 100,
		},
		model.Developer{
			AbstractEmployee: model.AbstractEmployee{
				ID:       3,
				Fullname: "Vae",
				Age:      22,
				PhoneNo:  "0987654323",
				Email:    "vae@sample.com",
				BasePay:  600.0,
			},
			Overtime: 1,
		},
	})

	avgSalary := 0.0

	fAvgSalary := ops.Avg[float64]()
	employeeList.Walk(func(value model.Employee) {
		avgSalary = fAvgSalary(value.Salary())
	})

	loserList := employeeList.Filter(func(value model.Employee) bool {
		return value.Salary() < avgSalary
	})

	fmt.Println("Losers: ============================================")
	loserList.Walk(func(value model.Employee) {
		fmt.Println(value.String())
	})
}
