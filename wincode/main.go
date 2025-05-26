package wincode

import "fmt"

type Data struct {
	count int
	char  byte
}

func Encode(str string) []byte {
	fmt.Printf("Encoding string: %v", str)
	result := []byte(str)
	fmt.Printf("Before: %v : %v\n", str, result)
	data := []Data{}
	for i, b := range result {
		if i == 0 || data[len(data)-1].char != b {
			data = append(data,
				Data{
					count: 1,
					char:  b,
				},
			)
		} else {
			d := data[len(data)-1]
			fmt.Printf("Before obj: %v\n", d)
			data[len(data)-1].count += 1
			fmt.Printf("After obj: %v\n", d)
		}
	}

	newStr := ""
	for _, d := range data {
		newStr += fmt.Sprint(d.count, string(d.char))
	}
	newByte := []byte(newStr)

	fmt.Printf("After: %v : %v\n", newStr, newByte)

	return newByte
}

func Decode(arr []byte) string {
	result := string(arr)
	return result
}
