package main

import (
	"fmt"
	"math"
)

func houseLoan() {
	var (
		yearR = 6.37 / 100
		count = 360
		total = float64(1_000_000)
	)

	dengE(total, yearR, count)
}

func dengE(total, yearL float64, count int) {
	fmt.Println("等额本息计算器")
	monthL := yearL / 12
	fmt.Printf("贷款金额：%.2f(元)\t年利率：%.2f%%\t月利率：%.4f%%\t贷款期数：%d(月)\n", total, yearL*100, monthL*100, count)

	// 每月还款额=[贷款本金×月利率×（1+月利率）^还款月数]÷[（1+月利率）^还款月数－1]，
	monthPay := total * monthL * math.Pow((1+monthL), float64(count)) / (math.Pow((1+monthL), float64(count)) - 1)

	//	for i := 1; i <= count; i++ {
	for i := 1; total > 990_000; i++ {
		lx := total * monthL
		bj := monthPay - lx
		total -= bj
		fmt.Printf("第%d个月: 月供：%.4f (本金%.3f + 利息%.3f)\t 剩余贷款: %.3f\t已支付金额: %.3f\n",
			i, monthPay, bj, lx, total, float64(i)*monthPay)
	}

}
