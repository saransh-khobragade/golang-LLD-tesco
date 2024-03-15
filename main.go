package main

import "fmt"

type Employee struct {
	EmployeeName string
	Shifts       []Shift
}
type Shift struct {
	ShiftName string
	Start     int
	End       int
}

func (emp *Employee) UpdateShift(shiftname string, start int, end int) {
	emp.Shifts = append(emp.Shifts, Shift{
		ShiftName: shiftname,
		Start:     start,
		End:       end,
	})
}

func RegisterEmployee(s string) Employee {
	emp := Employee{
		EmployeeName: s,
		Shifts:       []Shift{},
	}
	return emp
}

func main() {

	emp := RegisterEmployee("Mikesh")
	emp.UpdateShift("Bakery", 10, 12)
	emp.UpdateShift("Checkout", 12, 14)
	emp.UpdateShift("Dairy", 14, 19)
	fmt.Println(emp.Shifts)
}
