package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.1" )

func RwTYzYOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3zgtW2rbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9VzE7vBkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3oUMLc1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QAyUFgJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 03KqX5LhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VpYkONALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWVvP4fTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a2C2kQiUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uNLi0zjUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LOkL1AZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6WmecNTuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func md2LANpjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BCMgoIHMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LD9LRoFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func warEGojsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bnshiEbUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yNyrsg1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pXAPtzeNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDwyRfP8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I0Jj0jOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUza3iRIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sWt8UjnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5YNaTXgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WOt0zPVqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7TmW0nSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hgdyADnoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GzdNZdDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7ZHAG6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkVOPg6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2SXXHFXZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aCwqYATgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k6v8uLkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s7xVObIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p50T6pjzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qAJUuSBRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1RgMXYjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nxkkz99iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VkQZ7kGXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TXfHuqPYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rZPAe0hjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ighCLXb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nJ8M4WPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eh1XYP2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2q7YrCZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ImQVSds2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8o7PsVKOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTs19YBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vOhRPXEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJsUVtLIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jnrcQBGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lPghTYmeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HFhGdCdKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GbIMRB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zM1WVDgkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q03oxDesWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XYYo4CRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRnmuImWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xx5WauXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aVPrNWRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9DFTqEJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EChhLhTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F2sCP31DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XcKeyrfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GrLJ92s0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XSrfYPbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QNP9CAEKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GzlR1izRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PhdBwrF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FnSiZCQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VkiONXkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ih1qNifkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rZymUK5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f1GZz5POWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pmMSUDw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8TqerfJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lQ0zblYMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYUskMrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FO2zJXFjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8CSLvFq1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKV2up8BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SaWdDKz0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wRJsWyyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKyRxdOHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZBRgtgHnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func icyukLZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HfzGOwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n2NkS8WxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zZRzPoTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S38lTs7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EtFHFZKlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OisjBR3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EcyZkY27Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FhiMeBQsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yrz1QcphWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ZcGFXVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8liay0w8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c3tz6RE3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 82YIpGEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1vHR78GyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YusL7IdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MNKvM7E4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W4WB53PbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XsBH0HXbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8h8nivhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mYomzsCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ACPjOS0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PN12JR2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HhIf1xuwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRzrmlLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0BO1dv34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzRZ6zpYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZFdwOteAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZdVqqjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s9YmDoQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x2xFgjHWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hdYASMSHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cQzwJvKvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q5drmuKRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7dRD2nJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lkqX9CfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iR5olXMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GYux8ddCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zyFDUN3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFYup8dQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6flg79TMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZeQdUwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YiP2ob8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfLFb7epWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a9wVVxsfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UAZyDsMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cl5OITvyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9lddF8GUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KHva0zwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gIjkYZznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b2XZk7JAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 15w3BbtHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGC2DRawWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vnEa53QtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qW2OYtrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v61UtnagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PVO1lOcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQxAuM62Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v7cn6Ss5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aWC3UNDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZ4OwidXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func onYZf5sPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oI61Qte4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qwXYjUPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iqwIoh5KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JdiEspERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UsDGnZcNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lt4VH0ylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cRXNTbK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UNKLWGCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xCAxRUSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1i4Kq2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bBBGfZmpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ysSVnNZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HYjbNcgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hQGhgE8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H1cIRQ3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t5zk7U2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qsGyhYtbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjOTey5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WiwltSQrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xp2wf575Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aMoIaKKrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 46EXiy5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yr7cZ9wMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iegS1oJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j89BXV8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WtFOv0KOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Xocg6FxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a7UY1Wf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SiecZBmtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2oqkcqoKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6YTTwYlJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cf0FHBnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B721OkEWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQA8UJgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z9AC2WjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vk590f9ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 66lWdhHBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4avx809PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N3ViOFapWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbHUugfrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MB3lREMBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8sYAn0DJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bgunFVRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DlOxsrUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uBLhU0BPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OL0AQnzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L8cAlGLoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gnA2C2zMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gvU5nll8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UyiWbiq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7DKfK6GZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZ5I71UFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0siu70lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oinu3hgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Po9yDFW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WzMB8nRyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TdGHy4APWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TvBIhNPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 30htz8quWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func esTtB1T9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sppiv2F1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9atJ3QTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AkTdwQs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lHZdvkQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SJu2S95SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TiyiQxKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mxld5satWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gGu0Q21uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6MWBIbt4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7zYDdZWiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SDToBPLQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SnNnYUWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NadfyNwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uCSVlewtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BFkW8LzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQ4YWszZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BRKvzqiLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cg9Rlbc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U8DfbKbYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kRPx7yKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q8gsD4CIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJ1ZKziqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4sxKzKLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8srRMoxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9D5ABi32Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jSzfTe7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hh5hduwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WiwSuS1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yv7Pr1KhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbv0Hs5AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vkY25syYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BY2i09hEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EwSMCYqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SMd5ub1AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7wJj896HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YDklHD1FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YF9UAPWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w0vNCrk1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RvHjNpd6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MyfpoC4sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8JVc9fkkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P3R27wijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ssBk2xK4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VUXCs3TPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lnfI4nC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gZDl0W3dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zy0AIq0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I8Heqz7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5nSVePmkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aWXwHyHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQkgCPilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwB0M48GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xT0NLDtFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lb09xTlBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9uUcA9bTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RXEFwyovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l9vuU38mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M7yGtmH2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LFKDxMNdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ERXDFjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rIEIRJQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMnVKmpNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aSyxdLTWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W4pO4be4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wqg53aHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m3fcgG5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WcX8J7YBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FEpdTHEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8PvScB4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1Lhor8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hvzNVZ6AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func neCqy1inWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s7Wo8yhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 53hgxLccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func erXhxUxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3cQE2esZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MV04aaMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gntEkQM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func usqGUb6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2AYOV6IzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hDJDRQKqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5b8q0XMqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ze6anydTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jLv9BwVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CUEW0RxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6UG9miavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YwoF5UVqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func drRFX2bXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y98jIAVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1gWusv6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LFPbB2a4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dASSs3KIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0nAo5xjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7449reWJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q3CU4q2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func up2QwxVtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qSJzMBX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eRPHgaxNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func koVYHc8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YWmWupBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YbK3o0SeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HBl28EF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mxntsz00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L3bmOV8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IXo2q7HnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o3qGVSzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNgfiOtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UWgjNAwPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vYKAS7nfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PYqHcy2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func voGXwyHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SxFm6m5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1B6kS2bCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func noZEzK1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fwZCbiJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bjIW7nEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQrsh72ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zhBCUcoVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6mm92FYMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NEgDj7Z0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func knqDRdPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func my2F70H9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 44msMD0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vgnO6fzUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UYknhSvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func suc0rm96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5LuRsiWtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OYovqRgXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WyWcLXHsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VvsrleEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wD5mBCx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6Ewktq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DUDu3IacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CKni4N6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WoifVNbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4rC2IRaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L8iY2gAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pqcncsr2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SGTFOtl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qqQdKyjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8TipYj3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QeagPY5HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cPv5SnnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKilqcNkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U80CF0WHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6zLNVrdzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vf5krECfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H8sv8PAOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TztksIaqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RgrFkRx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hh1M5KV9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XHQ7FmeqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0cBOoFt8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUIp1cJAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wp0LFvFQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BIHojqA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LQvKJ3orWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OJM0ffD5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOoPk7gPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kyoB1W3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ATg4uD1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qmMrmxHnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTQ0I0oTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AujWwK2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WUWfBEauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jexgSHoYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2tg4OAohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Pz9945uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWjZSmAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PdEztdTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DoAILLD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Jc7fkCxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hHpp5mQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UEmcf9zlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vVqGT6W7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lvBHwfUwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J1Z51TClWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SoQbyfqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZkxQxuzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pwhH0SrQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGBFe2UsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HUt0ZTgaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gDbi5n8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YENERIQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pmDSy2NKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lIpPamAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vk5QdNSCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H1qpOBsvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mrpU5EhzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v5zdeUvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EkjbraTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HVvYTgUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 29fknZASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T5hR2GhyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lTrvYgmLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X49DqdwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYn01pyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFyO9c8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4QYPHIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zoyG2YjiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JpnQUI0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3PUHuoUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rITKNhgMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1XiipdGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T1z8jmLsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mWYFHs2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XIHrSLo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4hL1tUr7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WmKtNwXBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZdLB1HsKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3uiJdPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z3AT20LkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vVeoYbVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJz2SqXEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3uP5ODVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRW1thQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MPkxXLZcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nLdFHv8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0GI4yPkwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uzKGSmB1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6aAXA3TDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S4NtOXk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rnvgTPr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9QZ2X9alWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zYTTqRFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N47gZ0UIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DcWT7xMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8vwisH7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yXVKGKQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LQb8WiowWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0tWouliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 89fdbiqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PWppa2sJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gTr9MMXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wvh8RwfVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tPfOa6KnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NCGnenONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nQkLbShIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func moijXHatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func huenqoSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Om08Tq4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7bZHsuhZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 69Wf8vWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IdNZF4cCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w6GM2RWBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0LEixgZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RJhRDxV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3eZQ92Q9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4GRaYpipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c5UurfGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YuaMloj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6iTusjpOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUfckcY1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IAO5hGZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Le9Mr9NCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OgQKcgnTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tQycoBUZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzluZLb4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8t43LNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MgjCmse7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mYwB8EkCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Ixi7OZ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSxtlqXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ojg4HgJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XIiZkeTXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W6krcFUsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5V55MoKxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0jRsTYNRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gx4tFwwZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T4Ko1GUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GA0Qwuh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2U8tggTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RPHT6TknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aia0bbDKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oOmgIguOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rSwJoTDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JQqeQY3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZV9kaarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fyQ34fuGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ITGu0ZH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fgT1su0ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O3rkMfFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SJG2bDBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7vVjPTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KmxljcCyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8UbnnDFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twJFchpTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AaPiy2fUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bdUFETlaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJEfYQQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eCTN327KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JGoj4ukNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VP8RD5NOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mbecyXVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSsXdYvAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FvAMz3u6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z8iBQdAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X9HzC7BLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1coEtmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wwiB4M9vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CRD7wKVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HH2Jq5yoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJ5lPWftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tmpDCAaaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wlYSaY0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ReRwV1pgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZgVgcBHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EeXESvbvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMTxldBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6dAcseVqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cFMYnfP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e4jOnfJoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2T4JGSXVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rg2QGm0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fdLDSmLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZWQsBGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9mCAdxuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2wyZ2A9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XgctxP3kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QQDKLQ9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YvOkb6ymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MWDmDrgLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZcQLXdLxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D3xEzslPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OSIGULinWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jsaI96yKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fBKLaLUOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QsbZC7GdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jSVryGmhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xykuTdqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6qRVqi9IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kLHzUBemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7HH0ovh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3wzj4rN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ph0UgwMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uUOf4LUxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func icIwiAiQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xOO9MN5lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9alUPmW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ytBepwpGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p5nPrcRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZ7IqxSTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qxsXAKsjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zf5RMOclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fkUJiWknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PEZS13FyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func te9UX5tNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0woq7AINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TmC3T9qmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8RGYst5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MIpTC7qNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5n59bcloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sm2NGcS0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BbiSt9qAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8O01oqsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iDApEOW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AfJKEur0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n1qQtOGVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ECTdPKykWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nlY2l1T9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8v6mTHlNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lin0m7YOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JBfViBoXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ii4PNPeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OPOiZ06yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r3Wwm4ReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J32dTmb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xdiqIt02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sAQmj6WgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y4RrKTdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IAxHGUjPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wmf6dUAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Rn6An3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1SmqUfRaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EeUpK8RVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iby8VTx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d32Z3o70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kib4h7vnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SW55VIfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Mp5JyVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AilBySX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k2xuw62RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b017s0ZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BV7gTAjgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WzoGcdTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FZa8dMQAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func neLeCPdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SIrCDRtgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X2R0O2vJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func os45V2hLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1lSVpRwdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FfpLW2nBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rNUAMtzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func peNUs667Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y3SCs1xaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xqnYgvxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C1k5H1jqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lgr55kdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7gjra9LgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func euqNtFchWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QCJNR4atWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4jFPlvZoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJsZ8mRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQBf1zdYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IxwnK1IsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GL5Ja9OvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGTA5ZK6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JNSTpBXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jxRzjPzHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func od6X6STXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q3vDQVRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eOWcccUsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7PmN5sI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y4lhjH11Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYt6TawzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7kzr420wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kocHkFXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vVBQQdynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sMsxuTC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GiofTR2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 62WsEgdYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36kbhSabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2iIIHKYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lwcblx7MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FT51HvA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UezpNOXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDQLZJruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x5JKscwbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lnDXo0GVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cq9WRIchWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NgvFrgXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDuMhBIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KNyTT9qKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XiNAGi4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QLqb4mR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x9DSsBdrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pa04ihc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bW2tJUOlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ao5UhUq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nc8vtC1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VV7jTvdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BSiN1jXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W3EdbntFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PMLwQGavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4u9tlCuCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2E9xqPTXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VBftmINmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kUuadPUPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XEaxUfSaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 250KfzmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KzZdFxO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4M8kGx5YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kPuUOZtYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SgRghQsaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bzlkdRyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fKmCqlRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FpIV13dPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fl6Obl9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iS6XyIB5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GmM9xiY9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nf32mgtHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1CCliQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A0hdkCJoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eWJO8nioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1dR3AJSdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PKmC0q1zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func guXCmUBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4cKBY65OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twG3JmZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xWS9DHMPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bVpUsCRhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TuPlyePmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TuvGUS2GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7KjvyzOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQS6iVnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sr8fxCQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qCOlswyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4LpsM6UfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mtuk9IhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bbjke2CJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9q4e0jh7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uu2671GJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p82wpTzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fHwZ9jEBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GjK5ZA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qsDgz3XxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P3LsP4BxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tt8VnaOIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJZDWKdMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KNMjoGNhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DM1pnCeWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FM4AevI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hyWyoawUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PiMwLszyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rcdMJQhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lflCsjaHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hBEGIf1yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OsPjtGROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aWjB76fSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sHy0gHVwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ShsX1DaBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WVHlApt0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d8R16o4mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func woR96fZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DU3PwO4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hh7xXMALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S8fQWESUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OoSizUmEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 96345p74Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ZWVldOkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TGsdgCKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E5QjZtz6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7vZcfaW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PpdHmgOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FDnC0z1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func swt2sLt9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JpDCClVsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZ8ySSV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d2jzAVbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bW1vL5zxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zAJ75eSxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QAwG3Hq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ilr5tsnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jqG1rvNdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLUQa1VPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YDj1qpSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3lG2nIWWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lsJwcoISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SdWh1bRcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GHisXbvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sPFNpmhZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LeGu1fbBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71OpCONkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pvbZYNexWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0bIGKGigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v9586AVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWg5JoJSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hUl8p17SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJ4e6ZfSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZaZNbH3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NfzThZPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHu5SkEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T59znSVGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jai1dtCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iICx93W9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KXDwEHr2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kn16N2LhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NrRh8VDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sOSH5jp7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IfA7qatwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HMlLcAnSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9s55MdTVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7yknbOAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L2Ss2XJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GoNJljfjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 50dtcCOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OdYlYDYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AcYuXabOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7blxDK6pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UY65jHIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OsP1J49sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71PKSz2YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 19dbEnujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1dxOP3bRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9XjrG2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2c0C7WfXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F3JPkdkOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cvubll1RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YzqlspuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6uoqXp20Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TXRwhkUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LMITecuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tobxgI0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qJS5csAiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y6Shb64oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuu8NN9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b6Wo6kwaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjD2MoSxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DEGzfOtuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VrNYJxE7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GtHEeEuiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FAkjdVppWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GgwaZRCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wkz50aQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vgzMVWraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nDbMB4MpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sSJKai5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3osAASkvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LK1p6SXgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PaJRxmyyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cWH484UMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ol62SMsBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a1XKkBEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KU4Hr2g1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vNHuGKRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cy9F571WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PiHRB3l7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9U2imTkCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pER9oP18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nz3LuH68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SfFB6teHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4hnkU984Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mWjkPi8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVTbwdWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xIgDsc1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3f2XVv8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wDmVyZpGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LRxQZ56iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UjNyfIyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQLjX8EbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jvd4tsG6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S8GORa3XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2U8GYukEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func okM2g63bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pL6ECyYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ma5yGRH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RqeiZlavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P184e97ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5INOhxrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dc2izhCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5UzvD2fkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P2SRfMATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fR6noeyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gu2pOQmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vuujk9jFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2AGQAmkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6uKnYE1NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPidUyzUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pLXH9LA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 34fEeUJOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fTzayu1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HCsx9xVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3F7MljFuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LkGAcOuDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BKmr4zxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mF5bxFEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yd3U87XBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VUA763kaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nkS1lIUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQjJQAUZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1zJcuKwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LC36fXhUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tSl3pcPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G31SqC4TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHdcdBJxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MJNgr4hmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3RfWEQoxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WLwiQGDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhA4gUa2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4MxmToYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZpdjK63ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TWJXyF0hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LvQe0ATOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6LtJzF7KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9uWFtNaBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LurkdSHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func amJbodwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ptd1CdxJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func On1x3y72Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vio8yUQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4on5TnaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kygGuRPXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9krrW3K4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x0QD7T8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k531ehiCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IEpZiL3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tV9Kv8jUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GbxUpYXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tRxM5v1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t9LVMR8oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WHHj2ge2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func symvjyHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T1YUagtlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6rttuYU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ktCfiUMVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PChmMO7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JicYwXf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDe8tVhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A1SciaznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RqAUmZTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RzIu8xUZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8jke4yKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r0ysbTVsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9H3slcpRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LG0GKbmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nM7FR5UCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rttfXNZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zpoGOroUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KwoT0AEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H1jXkjoWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qRCnxoBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vR6pC8g7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4QcPGJdrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func njwVg2WQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hyNcLd7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q8bcE6iCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJ85Y6JoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkeb9hvQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u7jWrybdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k6MTz9pWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nBMJAKKfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h0tAJVSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lmnrquj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IjGDJTApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UEv5oselWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HP4aa3hmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jstewN9KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 16pl5QgyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N3Pht4eRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AsNS5sKNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TclX4IkfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c9dLiz6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hqvd47cVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KAHIVmlFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJSIysngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DbNcOEDjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TKaI4WI3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CO5NqdXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9gq9UUhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n31qfD2yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KrKiApV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bw5ZDmieWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iajTUTjOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func teNFnVuZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IVZ4SYFxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tDQA4za2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lUfZQcNpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RiwzZqpRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M5UspPWiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YE0M2OSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g7mbxXGeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zcHnsPj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sn7zx1KoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YBvKXISyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CALVPQn5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j0w0jCSZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9xMUcNlrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8R8JwcGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JqYrEsQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vv3LihpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ohBVCBHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bNU8FHZsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kTiiSwHGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7dFTdxHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qFIKBsIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ifT2e94VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YttBYADBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qaCWbOEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SiHxD3QSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PcHPkIQDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xFRRQANrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I43set30Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s3TnlMZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dsb7dJ8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9e7wwaf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p1xbgDILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNBd3OwNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

