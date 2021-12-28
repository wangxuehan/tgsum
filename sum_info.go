package main

import "sort"

func sumInfo(lenArray []int) (outStruct *sumStruct) {
	sort.Sort(sort.Reverse(sort.IntSlice(lenArray)))
	// 计算总长度 即产量
	var totalLength int
	for _, value := range lenArray {
		totalLength += value
	}
	// 计算N50
	totalLengthHalf := totalLength / 2
	var tmpLength int
	var N50 int
	for _, value := range lenArray {
		tmpLength += value
		if tmpLength >= totalLengthHalf {
			N50 = value
			break
		}
	}

	ReadsNum := len(lenArray)
	AverageLength := totalLength / ReadsNum
	LongestLength := lenArray[0]

	outStruct = &sumStruct{*sampleName, ReadsNum, totalLength,
		AverageLength, LongestLength, N50, lenArray,
	}
	return outStruct
	//fmt.Printf("%s\t%d\t%d\t%d\t%d\t%d\n", *sampleName, ReadsNum, totalLength, AverageLength, LongestLength, N50)
}
