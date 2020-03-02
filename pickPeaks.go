package main 

import "fmt"

type PosPeaks struct {
	Pos []int
	Peaks []int
}



func trimEndingPlateus(array []int) []int{
	end := len(array) - 1
	for i:=1;i<len(array)-1 ; i++{
	  if (array[len(array) -1 -i] ==array[end] ){
		end--
	  }else{
		break
	  }
	}
	return array[0:end+1]
  }
  
  func PickPeaks(array []int) PosPeaks{
  
	array = trimEndingPlateus(array)
	arrayPeaks := PosPeaks{Pos : []int{} , Peaks : []int{} }
	//arrayPeaks := PosPeaks{}
	for i:=1 ; i < len(array)-1; i++ {
	  if array[i] > array[i-1] && array[i] >= array[i+1]{
		if array[i] == array[i+1] && array[i] < array[i+2]{
		  
		} else {  
		arrayPeaks.Pos = append(arrayPeaks.Pos,i)
		arrayPeaks.Peaks = append(arrayPeaks.Peaks,array[i])
	  
		}
	  }
	}
	return arrayPeaks
  }
func main(){
	fmt.Println(PickPeaks([]int{1, 2, 5, 4, 3, 2, 3, 6, 4, 1, 2, 3, 3, 4, 5, 3, 2, 1, 2, 3, 5, 5, 4, 3}))
}


/*
func PickPeaks(array []int) PosPeaks {
  posPeaks := PosPeaks{Pos: []int{}, Peaks: []int{}}
  candidate := 0
  for i := 1; i < len(array); i++ {
    if array[i-1] < array[i] {
      candidate = i
    } else if array[i-1] > array[i] && candidate > 0 {
      posPeaks.Pos = append(posPeaks.Pos, candidate)
      posPeaks.Peaks = append(posPeaks.Peaks, array[candidate])
      candidate = 0
    }
  }
  return posPeaks
}

*/