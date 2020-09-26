package main

import (
	"log"
	"time"
)

// PseudoEncrypt 伪随机加密
func PseudoEncrypt(value int64) int64 {
	var l1, l2, r1, r2 int64

	// 通过取最高16位，最低16位将数分成两部份
	l1 = (value >> 32) & 0xffffffff
	r1 = value & 0xffffffff

	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ ((((1366 * r1 + 150889) % 714025) / 714025.0) * 32767)
		l1 = l2
		r1 = r2
	}

	// 经过奇数轮置换，r1存放的是高位，l1存放的是低位
	return ((r1 << 32) + l1)
}

func main() {
	log.Println(PseudoEncrypt(time.Now().UnixNano()))
}
