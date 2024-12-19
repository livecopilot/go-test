package study_map_reduce

import (
    "fmt"
    "math/rand"
    "sync"
    "testing"
    "time"
)

// 生成不同规模的测试数据
func generateTestDataWithSize(size int) []int {
    data := make([]int, size)
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i := 0; i < size; i++ {
        data[i] = r.Intn(1000)
    }
    return data
}

// 串行版本的Map实现
func serialMap(data []int, fn func(int) int) []int {
    result := make([]int, len(data))
    for i, v := range data {
        result[i] = fn(v)
    }
    return result
}

// 并行版本的Map实现
func parallelMap(data []int, fn func(int) int) []int {
    result := make([]int, len(data))
    var wg sync.WaitGroup
    
    for i, v := range data {
        wg.Add(1)
        go func(i, v int) {
            defer wg.Done()
            result[i] = fn(v)
        }(i, v)
    }
    wg.Wait()
    return result
}

// 基准测试：不同数据规模
func BenchmarkMapWithDifferentSizes(b *testing.B) {
    sizes := []int{100, 1000, 10000, 100000}
    
    for _, size := range sizes {
        data := generateTestDataWithSize(size)
        b.Run(fmt.Sprintf("Size-%d", size), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                Map(data, func(x int) int { return x * 2 })
            }
        })
    }
}

// 基准测试：串行vs并行
func BenchmarkMapSerialVsParallel(b *testing.B) {
    data := generateTestDataWithSize(10000)
    fn := func(x int) int { return x * 2 }

    b.Run("Serial", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            serialMap(data, fn)
        }
    })

    b.Run("Parallel", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            parallelMap(data, fn)
        }
    })
}

// 基准测试：不同计算复杂度
func BenchmarkMapWithDifferentComplexity(b *testing.B) {
    data := generateTestDataWithSize(10000)
    
    b.Run("Simple-Multiply", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            Map(data, func(x int) int { return x * 2 })
        }
    })

    b.Run("Complex-Power", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            Map(data, func(x int) int {
                result := x
                for i := 0; i < 100; i++ {
                    result = (result * result) % 997
                }
                return result
            })
        }
    })
}

// 基准测试：Reduce性能
func BenchmarkReduceWithDifferentSizes(b *testing.B) {
    sizes := []int{100, 1000, 10000, 100000}
    
    for _, size := range sizes {
        data := generateTestDataWithSize(size)
        b.Run(fmt.Sprintf("Size-%d", size), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                Reduce(data, func(acc, x int) int { return acc + x }, 0)
            }
        })
    }
}

// 基准测试：组合操作 (Map + Reduce)
func BenchmarkMapReduceCombination(b *testing.B) {
    data := generateTestDataWithSize(10000)
    
    b.Run("Separate", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            mapped := Map(data, func(x int) int { return x * 2 })
            Reduce(mapped, func(acc, x int) int { return acc + x }, 0)
        }
    })

    b.Run("Combined", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            Reduce(data, func(acc, x int) int { return acc + x*2 }, 0)
        }
    })
}
