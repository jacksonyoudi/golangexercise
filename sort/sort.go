package sort

// Sorter接口实现3个方法

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// Sort函数参数是Sorter类型
func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

// Convenience types for common cases
type IntArray []int

// slice 实现接口方法
func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type StringArray []string

func (p StringArray) Len() int           { return len(p) }
func (p StringArray) Less(i, j int) bool { return p[i] < p[j] }
func (p StringArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// 定义一些方法，调用Sort方法
// Convenience wrappers for common cases
func SortInts(a []int)       { Sort(IntArray(a)) }
func SortStrings(a []string) { Sort(StringArray(a)) }

func IntsAreSorted(a []int) bool       { return IsSorted(IntArray(a)) }
func StringsAreSorted(a []string) bool { return IsSorted(StringArray(a)) }
