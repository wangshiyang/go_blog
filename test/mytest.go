package test

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
)

func main(){
	start := time.Now()
	dictionary := make(map[int] string, 8000000)
	for i:=0; i<8000000;i++ {
		dictionary[i] = "hello" + strconv.Itoa(i);
	}
	duration := time.Since(start)
	fmt.Println("time used:", duration)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	start = time.Now()

	dictionary1 := make(map[int] string, 1000)
	for i:=0; i<1000;i++ {
		dictionary1[i] = GetRandomString(20)
	}

	duration1 := time.Since(start)

	fmt.Println("time1 used:", duration1)
	fmt.Println("----------", dictionary1[r.Intn(10000000)])
}

//生成随机字符串
func GetRandomString(lengths int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result [20]byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < lengths; i++ {
		if i == 10 {
			continue
		}
		result[i] = bytes[r.Intn(len(bytes))]
	}
	return string(result[:])
}
