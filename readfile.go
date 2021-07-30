//предназначен для чтения данных из файла, возвращает [слайс]float64/int
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

//функция для дробных значений
func GetFloat(fileNameDotTxt string) ([]float64, error)  {
   var numbers []float64
file, err:= os.Open(fileNameDotTxt)
if err!=nil{
	return nil, err
}

scanner:=bufio.NewScanner(file)
for scanner.Scan(){
	number, err := strconv.ParseFloat(scanner.Text(),64)
	if err!=nil{
		return nil, err
	}
	numbers = append(numbers, number)

}
err = file.Close()
if err!=nil{
	return nil, err
}

if scanner.Err()!=nil{
	return nil, scanner.Err()
}
return numbers,nil

}

//функция для целочисленных значений
func GetInt(fileNameDotTxt string) ([]int, error)  {
	var numbers []int
	file, err:= os.Open(fileNameDotTxt)
	if err!=nil{
		return nil, err
	}

	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		number, err := strconv.Atoi(scanner.Text())
		if err!=nil{
			return nil, err
		}
		numbers = append(numbers, number)

	}
	err = file.Close()
	if err!=nil{
		return nil, err
	}

	if scanner.Err()!=nil{
		return nil, scanner.Err()
	}
	return numbers,nil

}