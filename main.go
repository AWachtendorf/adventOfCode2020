package main

import (
	. "github.com/loslem/adventOfCode2020/adv1"
 	. "github.com/loslem/adventOfCode2020/adv2"
	)

func main(){
	v := NewValues()
	v.AddValues()
	v.CompareValues()

	d := &DataRepository{} //create repo
	d.ReadAllData()        //fill with data
	d.ProcessCheckTask1()  //process for task 1
	d.ProcessCheckTask2()  //process for task 2

}
