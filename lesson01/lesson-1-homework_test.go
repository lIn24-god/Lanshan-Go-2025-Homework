package main

import "testing"

// 被测函数为lesson-1-homework中的阶乘与平均值函数
func TestFact(t *testing.T) {
	result := Fact(5)
	if result != 120 {
		t.Errorf("Fact(5) : 期望: %d , 得到: %d", 120, result)
	}
}

func TestAverage(t *testing.T) {
	result := Average(50, 10)
	if result != 5 {
		t.Errorf("Average(50,10) : 期望: %d , 得到: %f", 5, result)
	}
}
