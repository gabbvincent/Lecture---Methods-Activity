package main

import "fmt"

type Salary struct {
  salary float64
}

func (s *Salary) SetAnnualSalary(salary float64) {
  s.salary = salary
}

func (s *Salary) GetAnnualSalary() float64{
  return s.salary
}

func (s *Salary) Print(){
  fmt.Println ("$", s.salary)
}

func main() {

  var v Salary 

  v.SetAnnualSalary(100.59)
  v.GetAnnualSalary()
  v.Print()

}