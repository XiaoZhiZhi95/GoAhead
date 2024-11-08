package main

import (
	"fmt"
)

type SZ [][]int

func (s SZ) Print() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d  ", s[i][j])
		}
		fmt.Println("")
	}
}

func (s SZ) NeedFillCount() int {
	count := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s[i][j] == 0 {
				count++
			}
		}
	}

	return count
}

type SZGA struct {
	sz         SZ
	cellRecord map[int]map[int]map[int]bool // 记录每个格子可能的值
	count      int                          // 记录还有多少格子没有填充
}

func InitSZGA(s SZ) *SZGA {
	if s == nil {
		return nil
	}

	ret := &SZGA{
		sz:         s,
		cellRecord: nil,
		count:      s.NeedFillCount(),
	}

	return ret
}

func main() {
	// originShudu := SZ{
	// 	{0, 8, 0, 1, 9, 0, 0, 0, 7},
	// 	{7, 0, 0, 0, 0, 0, 5, 0, 0},
	// 	{0, 0, 0, 7, 0, 3, 6, 0, 8},
	// 	{0, 7, 4, 6, 0, 0, 0, 0, 0},
	// 	{3, 6, 0, 8, 0, 4, 0, 5, 1},
	// 	{0, 0, 0, 0, 0, 1, 4, 3, 0},
	// 	{2, 0, 7, 5, 0, 9, 0, 0, 0},
	// 	{0, 0, 8, 0, 0, 0, 0, 0, 3},
	// 	{9, 0, 0, 0, 4, 7, 0, 2, 0},
	// }
	//
	// originShudu.Print()
	//
	// fmt.Println("--------------------")
	//
	// szGa := InitSZGA(originShudu)
	// szGa.FillAll()
	// szGa.sz.Print()

	originShudu2 := SZ{
		{0, 0, 3, 5, 0, 0, 0, 0, 0},
		{0, 0, 0, 3, 2, 0, 1, 0, 9},
		{0, 9, 0, 0, 4, 0, 3, 2, 0},
		{8, 0, 0, 0, 3, 6, 2, 0, 0},
		{0, 1, 4, 0, 0, 0, 6, 8, 0},
		{0, 0, 2, 4, 7, 0, 0, 0, 3},
		{0, 6, 1, 0, 8, 0, 0, 4, 0},
		{2, 0, 5, 0, 1, 4, 0, 0, 0},
		{0, 0, 0, 0, 0, 2, 8, 0, 0},
	}

	originShudu2.Print()

	fmt.Println("--------------------")

	szGa2 := InitSZGA(originShudu2)
	szGa2.FillAll()
	szGa2.sz.Print()
}

func initMayBeMap() map[int]bool {
	ret := make(map[int]bool)
	for i := 1; i <= 9; i++ {
		ret[i] = true
	}
	return ret
}

// ExcludeSameRowAndColumn 排除相同行、列的数字
func (s *SZGA) ExcludeSameRowAndColumn(a, b int) {
	for i := 0; i < 9; i++ {
		if _, ok := s.cellRecord[a][b][s.sz[a][i]]; ok {
			delete(s.cellRecord[a][b], s.sz[a][i])
		}
		if _, ok := s.cellRecord[a][b][s.sz[i][b]]; ok {
			delete(s.cellRecord[a][b], s.sz[i][b])
		}
	}
}

func SudokuIndex(a int) (int, int) {
	if a >= 0 && a <= 2 {
		return 0, 2
	}
	if a >= 3 && a <= 5 {
		return 3, 5
	}
	if a >= 6 && a <= 8 {
		return 6, 8
	}

	return 0, 0
}

// ExcludeSameSudoku 排除同一个九宫格的数字
func (s *SZGA) ExcludeSameSudoku(a, b int) {
	rowBegin, rowEnd := SudokuIndex(a)
	colBegin, colEnd := SudokuIndex(b)
	for i := rowBegin; i <= rowEnd; i++ {
		for j := colBegin; j <= colEnd; j++ {
			if a == i && b == j {
				continue
			}
			if _, ok := s.cellRecord[a][b][s.sz[i][j]]; ok {
				delete(s.cellRecord[a][b], s.sz[i][j])
			}
		}
	}
}

// SudokuRemoveMaybeNum 某个格子对应都九宫格内的其他格子，移除某个可能的数字
func (s *SZGA) SudokuRemoveMaybeNum(a, b, removeNum int) {
	rowBegin, rowEnd := SudokuIndex(a)
	colBegin, colEnd := SudokuIndex(b)
	for i := rowBegin; i <= rowEnd; i++ {
		for j := colBegin; j <= colEnd; j++ {
			if _, ok := s.cellRecord[i][j][removeNum]; ok {
				delete(s.cellRecord[i][j], removeNum)
			}
		}
	}
}

func (s *SZGA) InitCellRecordMap() {
	s.cellRecord = make(map[int]map[int]map[int]bool)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.sz[i][j] != 0 {
				continue
			}
			// 初始化
			if _, ok := s.cellRecord[i]; !ok {
				s.cellRecord[i] = make(map[int]map[int]bool)
			}
			if _, ok := s.cellRecord[i][j]; !ok {
				s.cellRecord[i][j] = initMayBeMap()
			}
		}
	}
}

func (s *SZGA) FillAll() {
	s.InitCellRecordMap()
	for {
		if s.count == 0 {
			break
		}
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				// 如果有数字，则不需要再计算
				if s.sz[i][j] != 0 {
					continue
				}

				// 去除相同行、列的数字
				s.ExcludeSameRowAndColumn(i, j)
				s.ExcludeSameSudoku(i, j)
				if len(s.cellRecord[i][j]) == 1 {
					num := 0
					for key, _ := range s.cellRecord[i][j] {
						num = key
					}
					s.SetCellValue(i, j, num)
					continue
				}

				// 看格子可能的数，在该九宫格内其他格子是否可能出现，如果都不出现，则说明这个格子就是这个数
				num := s.CellMayBeNumExistInOtherCell(i, j)
				if num != 0 {
					s.SetCellValue(i, j, num)
				}
			}
		}
	}

	return
}

func (s *SZGA) CellMayBeNumExistInOtherCell(a, b int) int {
	for num, _ := range s.cellRecord[a][b] {
		rowBegin, rowEnd := SudokuIndex(a)
		colBegin, colEnd := SudokuIndex(b)
		exist := false
		for i := rowBegin; i <= rowEnd; i++ {
			for j := colBegin; j <= colEnd; j++ {
				if i == a && j == b {
					continue
				}
				if _, ok := s.cellRecord[i][j][num]; ok {
					exist = true
					break
				}
			}
			if exist {
				break
			}
		}
		if !exist {
			return num
		}
	}

	return 0
}

// SetCellValue 给某个格子赋值
func (s *SZGA) SetCellValue(a, b, num int) {
	// 赋值
	s.sz[a][b] = num
	s.count--

	// 将九宫格内，格子可能的数字都相印去除
	s.SudokuRemoveMaybeNum(a, b, num)
}
