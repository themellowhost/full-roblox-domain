package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.7" )

func EoUNsRG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gXVrYQn4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kUsavsBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U3zEZliqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kAdpsIoCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sLlrGSMPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WuSczs0hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PVHpcW8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Zfx7nG6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ww37Tni2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fiQEEaxjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E6wiWtWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7qawP3ryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ixay8HigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zUhTQdDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0DoD37MEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xdN6JfwKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cqGrMztUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3FCCOlw2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 06KpPjOkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 91t0HIB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0aI0fju5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5iP6sju1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eY8zYFAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jva2QmKRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hcomI1UMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bpS44sA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OkzvPnNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3kWDkDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7V4zT6RsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LXYzu7DsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gHboLOLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QVRk1jbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jirD24R3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RkHxvkSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U1OgUl3lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xMRhgCzaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QByLZqV2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xodd8bLdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PIijpvHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k1s81GIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q28yK8KrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KgPSn7yNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0993vS6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zJc725ZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tWlmeL9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wgueW0McWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z2YcLcBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YgujO2avWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z8lH21K8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wyP48gLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1nq7HdPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3TB269pjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func md0JJu3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fgPk9IEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 53m8bA5HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8DQwjYwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ogtn1ezzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XOed6yukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S6JtB24cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g5Q9hsKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l1Zb8N8iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yw5eoTftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jHRrC1chWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ir3CofecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wzsHIN0hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lq6ABKGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ajmkzen9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjdvbbdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DH5kTvPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v1Hw5wqdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7V5XuAuTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZpJxP9JIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQitWZKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tal32FZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func reWWvDtOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tVsFimvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2fiN0iAgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F1gF2KmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TXAhswFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func prbbtQ3kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l3oeeQyaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJaTWbW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y9pqyBxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func axoLXgKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W7vm0z5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HKYBdO2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dyyjZyrcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func neu5XuJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3yIaJgiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CyhivDpRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nd8b1ES0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3jW0yh5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XCeW68iKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NVYvJ844Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OABcG9nzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TFZseL5nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fF3pL8naWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jc2p6Y57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AXlndUnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 165W3dYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5omvKkAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rus0KOanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E2vxcGqwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1LvO8xgMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhq9xLn8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GgSyyaJLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BXT2xgQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oxpcwnBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1XVwvMU2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0GxKygGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WbGEAh6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mv1ipsnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zm8fVoRdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XwJCjOuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cD1LpiMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tD4Q1IY8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ifIIweUzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gpw659ewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ott1RObCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3czqUKU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AD3z6XFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JuSXagJOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ea1ruphjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lr3JGeeVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func swhZeZTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UNUVPpfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f8KPJklAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cPfJW2z4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZaNMNCX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TzgqbC2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TzlGBQgMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r4qnUEXuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func amwKrDlYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y423kuGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BpT3gkhbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func auqrGnxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gm1Jd3pNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cZ8jso5dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UL4B0IU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wjGozIwZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MwtLoUeHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2RJIaPzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BAT80dYHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fe4VqhY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tVFnzWIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w6GRt3n5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gHr1McfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bF15CdhWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VMUqZwNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7JuksLS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eiX7D5RnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FrtZeNx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w14JvrPCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jSV9n3wsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4L9COWEjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RklSnO4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i7GpzTY1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4hIu5mDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s4UCbzhPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gOjyDpmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S0z8rrQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPpIsLxNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func erOzt9HIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5My6CYI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zRDq8PsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UvcGE5MXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WOWclyYMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6V9NeG2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FDTlxFudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XLFuP08hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yeCw8ItIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UfXySQeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Ggl1ChTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D9l4VgPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Z7xVuzaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMLm0t1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UWiuJnlbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z5Q7P3qcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1DeYFJjdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nyvf2JvxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GSnPJo1UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOtGpUMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DqRbSrXeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FtKeDuEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u6IwxMvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VJJNOdHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wwBZMf6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EDtcJXeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BYGPGcKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sRmK9a37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ofuz9v7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CiRB2p3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l3GeMMfoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jUeO3wdmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rL9MpeTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IGIoFYlLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PWrFCOG1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K8Un0rvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f6xX813gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fRDGeOHdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func isFLrvm5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qBTAC2R4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jt3PQIpsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XS3DaQHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5l58xe28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQlJpKz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zxuMJvcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func psxYfEeEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TJ39oEmHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gbpt2uuwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pefidkBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func me0Mzo10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pXONTtdsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pgVd1Rg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0mkp7XOkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VkTegyWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XGrqmhUjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MXUFTMDZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Uj5itGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GcBnKkY1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sWfzOBXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVjn6ZzuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zlw0SPDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mkgcw9JOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RnW2k2twWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SIFJ0i2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func td4EmvSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cPgoDIblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q9AHP1vGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jHRqnkdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NaxwwUKWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IBfWMb7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PRtiaqSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6jy0VW7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Shr7xI4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4zvgyloRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ACcNVfSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ClGi7lQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VHyieXbGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ItqxfLWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4irRkEP4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MJYsgmOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kFpnWxy9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JcTL8jn7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MR6LfJRdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IRyu8xd2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 23RA8Kn8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fVLs5TjGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DmFu6Mq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cnxIY3iyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8F6v4iYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZTZEQDqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func USWUbNjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zadbrsTyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LX94y1cbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pRausIplWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mshXCiWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZQvnrv3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rKQE6pFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3h4jhLLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aF4XbVLdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 26vm4UXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 70wGKJl6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1t77kTKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIs2CzJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TQ2lElayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R2NGUU6UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQen0zJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TOdLcaNpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ci054VRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D4hkogAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sDNWfxjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 03huPSuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xkOT6TLwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8vMZ69EjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ECSy31BHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S60DG78hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QGi8FY6YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X71H1PSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1JlHu0dbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wWnvTbXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uA6CNRzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZcgvkTTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZNPi7jWBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a7L0Dk26Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sv38Y7z2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qPJ3Qh6jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cb2jAj9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2StUkwxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RrSj9OZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func As4zkfqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bFFDeaINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func irEg4VUyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6qvM8PfSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 68p22GtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6yQEN0gWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rFJHTqXgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 58M3h5MFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7cpSzGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jcHTYOlKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mxZuBcvsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sIWVSadxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5GGIm2beWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EqAnQH1zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8TkqZgTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6eRWErYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2hiS2F0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JsPAQ7rzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lCkRtBGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpsGSY2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wLCx3DAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ediYn7xbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5N8vUWoXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vnSVvyBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QFof9gK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oR9gxR4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eXWiaNSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQTv6j1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuLkuXRsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNGQ5zOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SCeehE2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5xz7POHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O6Ct3Mc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OB4NsxvPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuY8wHwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func plFPLhp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PBDGzFMKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qlrDzbwdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9C1FpmyWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hp0LXq9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HYT7eFtEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vDoxdlrSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2F9m5uPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U91ssnOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZVsIwCGEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XexLDhI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oyvaZfqjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 70CmnVyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QYqAnTUFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EY1TL87TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJwmroIhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FEelRp7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jfv8ZjH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wog2fufSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eo2zcZLoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u7xnzo80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ixwm3ZjDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mxI05md7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func krW9OmW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dw1JNNHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C9bDPbRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aCKTEpLpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1QQarfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5D4RbvxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1SbcjpnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4YfOErloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ldNjCXtIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h7X2r2ONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XbL4db8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2AROCTm0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mGIp9FzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4UsvIuitWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iavvEXo1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDumcbsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjFhyGwoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mxbSL56vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UjfSq7RaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5OWSWyJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bCE9kXeRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4WIWnx6kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zDjTcfs2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qkgMqvF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hLx0nsBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LjP3rg7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MyhPToXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRVv0fozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XW8ux4rLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func avd82e1cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZ2p4aW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D1tpl3vQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mbn2b0uCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QJHKcVkQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOE2sdGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FS25IfULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ID534x2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Df9X5ysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YvopaTD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9SVfQGx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 211lacQYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FhaxLNJOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gGavV4SlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EbL9f2BoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pkeJa2pBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8LHGl2DKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yiWpQsg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yKeDX1FfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MXSlw8hSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mgGxTsVpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12t2LLjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xAe2G3lmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CalRUPSsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJ4zvIviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ry2Aici7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Or3ZARnzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ue9rgUrDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uMDws4X8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5aKrTUI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AC6XbvjoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j74Q72UfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJvOIT9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MqVkoeEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func skL5Qa7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XQBEQIGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iqUTrdYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n3x2IciWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OzCxKfsqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mKUhEJNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p3dzRxjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DCXvcpb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SJJXwTZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lgMvvzT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIiPCj62Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6kKqCewaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nf1gzbbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5H1hUgSAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Lgp4H6fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ZKdxURFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CsNIrHKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFWoWEMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xBANyfOIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YSzInW3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YLGSUoPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jRtP3AHMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNxPQ0eqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XP6q72TCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wB2Ce0zTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lADnpa27Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func of2eTmDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8VWHSipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JC3jzzkLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eu4R0sziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BQhOGrgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kC13F5wyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5B4be4E7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGLF6rw7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MGZEWthhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 758o0uXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0QIJUthWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xUUNmDMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2qVkLAedWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rvmr7s2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4vXC21UsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQdX9YysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CTPQzGHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HeANbbrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pmGrRG3lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VsmNo5C8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CAMbdL8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aoBwsiO1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M7FBBew8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yH24wzsfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4hnEqCKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ot3DjtXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d0UwjMmlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KB9w5ZbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RuMYTd9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2IF1cdDVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O5RqVaOIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2QGxSnq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EWCkcnqbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Fu5FNf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RqTWF1mcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7UwdM6RvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qGYRMaQxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ulLg0twIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HfUJrdh8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRjfaJFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0eOQSjAZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dbAEnT6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jzlah5DRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ETtbvjpJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oPAvwSyTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h9wKu961Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uiA5EtDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2oOfYMkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hYFLqF1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1ZNvuWJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jh6xbDiXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YvKzREHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R7FAq120Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kVnyOVMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func guIb9XlCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kuJ90fT4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TvrogZGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dkNV2um1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2ngSUMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T2DfGPQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FijPZK9fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hk7twQAvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xsnIecDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jtFSaTd6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q9uEwCCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REMsn3OgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q1jQyYguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cabQjHLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZPJR16n6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R90xm0mIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DPgprD11Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func axdQrWJmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Xq5ttxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wCr86zYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tyyVQ0tEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5mbhGMSfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ijdZkjArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFjyE8QSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O4AdPVMFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m0FjqN0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SgpIqfXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5D4csTDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func it4OUyo4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4SKTs7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vXoa3ZwbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BiaAio5vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cEfbXTrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QJKGGsrKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J6yzVQeWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tt4baYQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mABb0Q1WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jrDjIUv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xHFHsCXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yHrGMbIuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbMZkuISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4AgobHjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9BRcCtCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7PtfstGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ndtboWMxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WMRvaxmjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WCMLUjwsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oPbu9B1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func waDLcUuXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8zo7XlDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wrE7MooxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NO4iREnzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Go5EgUR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BxZAhufwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fxjnv5qOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbZqFc9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wFKHYzafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qFdAFM18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LU4gqNX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0pOGcoGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TvdzoXAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2DxX3mIqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F72gcNyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9dMhSFfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1svzsSjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3C0ysVIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zhBpnZ7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bVU8JPIhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BnnXtzUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w7fALP7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g0Q6C7brWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mrir5BJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uv3cHlhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5SGux9F0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjNE5a98Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3EMzLdHMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KS9KJqhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tWwF4XRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U1P4fYGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b3wUCc7eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qO2OYIG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XjNoFllyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6FPcL4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t1Mx8Z8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RnHhGIyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4cyzajDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jiRA46RdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JhXoIuDLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GKNNgd6YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9CYLCOPCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oRak0gRYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ugXcgjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9jEFZfciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y1BaMqQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nZRjC879Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L5qDliV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EUa5BuJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DbSlYk7QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3L4XfunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8AYGhXWFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3eC1JNbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qiGvm9MPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bBsWw9w8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4wnYIh6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rG7NED0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RtyfjA7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func axwJY4OWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vlWIzv4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qjTksyKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kQXsAwbVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 47MycG9FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdTuFivzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B28GUMf3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func urddbPJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ogQtKGumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FYZqXmhPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QFjNoC2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P5ISO28xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RgyZVKzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umXmJSmSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9XxG5cRYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQiZ1WgVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M7gg5pvIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2GuN7XNzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X1v7CXsuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYKsGisOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eI7ZbNdPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NwM0rkYVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ejkUT0J7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iWeHqEmXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZttVhH89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QuQtPz6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EudsaK6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PG5r2dzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kHRIv6rgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func giabzALJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v628RKjrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ievVx0f3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tErSHCqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4RHTLEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q1wc0rzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bzTbveCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M3AnFX8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XubBiY6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QIvvqSYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zstL81zEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bLTIh3ixWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2wLqcM3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LwpjJIbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qcv2nq0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SRuhZJsUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yzE2ikbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W0hWpkRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uT3xEBi6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rEsFBKpzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bs6fDvXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZHrX6T7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G12ZDzlvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pfBVJbDGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFPEwdRFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l71klpwbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ykAP5m4GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mKhyfz6gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLS12oupWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OOuLkyFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fICaKVq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ecQPImnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func emhlV6yOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tTlVhbawWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tlgkTZX4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iHBewLPjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ibd5Y7oMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ktNMPDq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FnvRweFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZ2gtjc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A4C3m2pbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3oSAQFfKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YlEeyimVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FIUsO4BPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iNGCKnseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OgM8ES5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func osRQfVVMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func chwNggbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VYFqs5cTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bv6ov1i0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fzuCsuTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HuSB8srHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func orbTd7oaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3eIAFpGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GccaZiJkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZFxl06TKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HEtDahMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x01WEgBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4U51EVHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjBEX5FMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UFde0QU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f53RaT8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KAjwV4b0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZPsJy0L2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FmEtYgTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pdXV9EgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ShVmsJwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1MR4yhCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HFTA7U0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lr77isigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8xmEwfygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3s7Qkzc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TBZEWqUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dN5KBOGAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMTmLILvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rm7zQARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TQheyIGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kuFpNHbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJIJkdhzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h24I4arbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8mGjA71HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SFuu3mGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3N3d52DWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ucBpMaCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJNT1GIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d0kyxIxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KCJc8rPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nnBkuXYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G4PeIcDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pi9UYI17Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdXoRZGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3VaaOEV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oFr92ifjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 87OYnKd9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x5D1W5KZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJrUMwU2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lt62L9M1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5TjoY82vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o8yHGk04Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RmenQrPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uA9m1YonWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s8RnATpgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4RYzEOWJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kT9zpKnnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lkyFyNkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2XVo87wRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 51IlIT0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x12KBgg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQuq7KXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sVDwjXMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ivmj2b6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z3S1kllCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EYHZloIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func efZ563OTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func otidYL1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PRmdltAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rwOTYzEuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v79xDcY8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rtsngUz0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jekLa633Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WEAdym0hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fKQWshj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ao9qXl31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fNl8UptaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rvtUU5tnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vVaUfJukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O7RzWnkFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D55L7LzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k6WY8wk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g35pFWaJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mTJm4ouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func inyhJ4m6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0dFX6KzUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 59hmsiyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWWngZFxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J9oyj8nmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vINoF8PVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mw0LpkdtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kKTJxBydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ORusU8j6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JOyUy8nVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GTczELB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iMPD16MqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vl1FOfFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gqYJamaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gQ58QIg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MGk6SglQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 64Z5SBOjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7NDQah2TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QNPPDxO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z0oiB5GRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w51ej6nFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 16iVNZj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1bw8uee6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fRtMF95FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVr3A5D1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cQ3up9wJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K86Ae95rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yotXcy0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W4Uk6azWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h5NfGdVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qLM3CNfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWvz4QKZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 27nXTv6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X97m2ygRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AOFf6LrXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fh8jnQkkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kmCuLJgVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1gRK0svpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYL3YEqyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Axh9TQVJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bw5x679vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s8NuEV0wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wl6w9QJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQw9Jh35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LupvvDOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tu0on96wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TCphthydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RKMlBe3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dwOG395nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Tlwox7KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgHHMBxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5DS2QdbVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z6fGkye7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GdVL1k4CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIDMUeNzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DXWKnhXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fpybmy02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func amVbDii3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3JRlPLOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LmvbkRLPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cuKf2O6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nwpr7pXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ITU1GeEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Fq3IjyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgvg0mZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wIKo5xq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ltkAqaKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kVX5qylOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IjjAEY8cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fiLX3QEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9yhkO5i8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BFi615b7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r5PvBEyzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WEo6ndkAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nrIM2lRcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BPSK87jJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pBHZRaLxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxFCfACbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N18jqmhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ag60QYp8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qe3kadDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2wyQIZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5azMHUUzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3uiOj3MaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A5gadGhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X2XpEhIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ZXlillPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7K55PMo3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HkQOEzQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yoB43jHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UsPIYS5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKs0n9D9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k6KMYnFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n51Okzw6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xLtl6qzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PEfS5uobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SVI7CuBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tNHRQ8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRcNnoDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eYJRswWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZT7oIBXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rhPEfvbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func migbCUaLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EDNhiR28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lDJhogacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thpUmX37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 59Ihc6YHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jrWxMVlEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thNsuLb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4oM5LjVEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WUUzsVPYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vYgcnJdiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lIrgtdwxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bbUibapuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oUr2tc5EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CcPliF8LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XLAWQD5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhWelRaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DeLRGg82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WmR3y5ZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wftccMrbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zDSyPjh7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvLHIz3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GPTXIL8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V07hybouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VyjYqytsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bKR2k5mMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YxYLX1FDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zfesabs0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uH2ba4I7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ciNlkio3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fE9Y0c7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8fM3wBv5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tmvsm7AUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V4l8NtgaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rE0QNMq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ogVS5zYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iDbB9v1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xUuRtqUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2y69QGjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vsFQouDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ViPoTaW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KBl78FKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wlPR5Y7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WVfRfLGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ICcgBpaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5CfFkPuKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Doe034A9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z1VuUBp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YJEF42oDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LvH4U7IxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jYFq7ntWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 633vh7wjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9UYBFZh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PfVD0Oe4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g0oPh6hBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lHRnU7jhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DJxTItZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nyIaVl46Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BDIiH9z6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7PuMHhcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LK5kQYxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1vpgMHozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x4eicjk2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ph9eWjiiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EN466Zl6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rvY0QZ99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rkvTt1KFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fs0FUyemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYN9fzcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func okJkLU64Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qf303f5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vxk0sr7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vOzEBFLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHULnAW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w2UwRhIUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r9EAGxXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XrfF7h3ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func az2bR5rDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cbD61UgYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zAjqxnO1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oIT6V7ZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tk1vcHr8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4hTDz76LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6B8OE5asWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t6vUDu0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EuwYsdr7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VV8VGRTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMKTdKjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QsFnOOadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GVXDU8LeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eyde2FmGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PSgcx6VbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UMhyhixgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mcEVFpfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pX0ZkW3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Id6yDsUbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LsVLoyjzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6OX4xWxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PBELZ7sNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xq4AX5rtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r538rg0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A0WaYRdQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TV4fYlBpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9y31GqmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func awHqCJUPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LZuvtzAxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gjqCyeWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKFTEUz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cEhkxntHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c1lNWZfSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1moUIfkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DyQmjFhvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oTmUVEMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJdiT1ASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5TVjTEqkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPlOzwEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mTYNTt2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X6dr7NnrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BdpU6qkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3rHe731SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tki6O6J8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fWcFjgdqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y6U9zh7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZQCUw4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x9IUnS3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Z52HpT2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hw1FzvDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4f2kVFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FvozvPKWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func smqoStzNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func weXMUfnqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func npUTXMxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1fZckiLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eBnxWdyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZuIfO8RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tvpgq7mXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZtH4iPEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ukl8AXiFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EXZSmWF2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tCdqeWV2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3g1TLvkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func popASDTAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5MIxX6UTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1PrcX1VkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FxwCrd8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LS3cZgcfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TtpPU83uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M1RYvm6kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k9SREQZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xzUBbSOtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5tBo9qD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XzKBitqeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rfywjemAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hNoJDt3YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VphaxzyrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTx9hgtIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JaDgSjEmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7gNRhrcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NrO1khf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hWkSgwqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XrxJC3BTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NtnSN6ypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qn8KbY3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Fp9imIUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMKlgjvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMtkGP8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LMJ3KGdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

