package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// routine concurency
func Concurency(list []int64, max int) []int64 {
	var (
		i            int
		listReturn   []int64
		listInLength []string

		j      = max
		length = len(list)
	)

	ch := make(chan bool)
	mtx := sync.Mutex{}
	if length > max {
		for j <= length && i < length {
			listInLength = append(listInLength, fmt.Sprintf("%d:%d", i, j))
			i += max
			j += max
			if j > length {
				j = length
			}
		}
		for i = 0; i < len(listInLength); i++ {
			spltInLength := strings.Split(listInLength[i], ":")
			// Routine to spawn process
			go func() {
				mtx.Lock()
				index, _ := strconv.Atoi(spltInLength[0])
				length1, _ := strconv.Atoi(spltInLength[1])
				listReturn = append(listReturn, list[index:length1]...)
				if len(listReturn) == length {
					ch <- true
				}
				mtx.Unlock()
			}()
		}
	} else {
		return list
	}

	select {
	case <-ch:
		return listReturn
	}
}

func ConcurencyV2(list []int64, max int) []int64 {
	var (
		i          int
		listReturn []int64

		j      = max
		length = len(list)
	)

	ch := make(chan bool)
	mtx := sync.Mutex{}
	if length > max {
		for {
			go func() {
				mtx.Lock()
				if i < j {
					listReturn = append(listReturn, list[i:j]...)
				}
				if len(listReturn) == length {
					ch <- true
				}
				i += max
				j += max
				if j > length {
					j = length
				}
				mtx.Unlock()
			}()
			select {
			case <-ch:
				return listReturn
			default:
				// DO NOTHING
			}
		}
	}
	return list

}

func ConcurencyV3(list []int64, max int) []int64 {
	var (
		i          int
		listReturn []int64

		j      = max
		length = len(list)
	)
	ch := make(chan []int)
	if length > max {
		go func() {
			for j <= length && i < length {
				ch <- []int{i, j}
				i += max
				j += max
				if j > length {
					j = length
				}
			}
			close(ch)
		}()
		for i := range ch {
			listReturn = append(listReturn, list[i[0]:i[1]]...)
		}
	} else {
		listReturn = list
	}

	return listReturn
}

func ConcurencyHTTP(list []int64, max int) []S {
	var (
		i, k       int
		listReturn []S

		j      = max
		length = len(list)
	)
	ch := make(chan struct {
		list []S
		i    int
		j    int
	})
	if length > max {
		for j <= length && i < length {
			go func(iarg, jarg int) {
				resp, err := http.Get("http://localhost:8080")
				if err != nil {
					log.Println(err)
					close(ch)
					return
				}
				body, _ := ioutil.ReadAll(resp.Body)
				defer resp.Body.Close()
				dataSend := struct {
					list []S
					i    int
					j    int
				}{
					i: iarg,
					j: jarg,
				}
				err = json.Unmarshal(body, &dataSend.list)
				if err != nil {
					log.Println(err)
				}
				time.Sleep(1 * time.Millisecond)
				ch <- dataSend
			}(i, j)
			i += max
			j += max
			if j > length {
				j = length
			}
			k++
		}
		for l := 0; l < k; l++ {
			in, ok := <-ch
			if !ok {
				return listReturn
			}
			listReturn = append(listReturn, in.list[in.i:in.j]...)
		}
		return listReturn
	}

	return listReturn
}

func NotConcurencyHTTP(list []int64, max int) []S {
	var (
		i          int
		listReturn []S

		j      = max
		length = len(list)
	)

	if length > max {
		for j <= length && i < length {
			resp, _ := http.Get("http://localhost:8080")
			body, _ := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			l := []S{}
			err := json.Unmarshal(body, &l)
			if err != nil {
				log.Println(err)
			}
			listReturn = append(listReturn, l[i:j]...)
			i += max
			j += max
			if j > length {
				j = length
			}
			time.Sleep(1 * time.Millisecond)
		}
	}
	return listReturn
}

func ConcurencyV4(list []int64, max int) []int64 {
	var (
		i, k       int
		listReturn []int64

		j      = max
		length = len(list)
	)
	ch := make(chan []int)
	if length > max {
		for j <= length && i < length {
			go func(iarg, jarg int) {
				time.Sleep(1 * time.Millisecond)
				ch <- []int{iarg, jarg}
			}(i, j)
			i += max
			j += max
			if j > length {
				j = length
			}
			k++
		}
		for l := 0; l < k; l++ {
			in := <-ch
			listReturn = append(listReturn, list[in[0]:in[1]]...)
		}
		return listReturn
	}

	return list
}

func ConcurencyWithSort(list []int64, max int) []int64 {
	list = ConcurencyV4(list, max)
	sort.Slice(list, func(i, j int) bool { return list[i] < list[j] })
	return list
}

func ConcurencyWithSort2(list []int64, max int) []int64 {
	var (
		i, k       int
		listReturn []int64

		j      = max
		length = len(list)
	)

	ch := make(chan []int)
	listReturnMap := make(map[int]int64)

	if length > max {
		for j <= length && i < length {
			go func(iarg, jarg int) {
				time.Sleep(1 * time.Millisecond)
				ch <- []int{iarg, jarg}
			}(i, j)
			i += max
			j += max
			if j > length {
				j = length
			}
			k++
		}
		for l := 0; l < k; l++ {
			in := <-ch
			for ; in[0] < in[1]; in[0]++ {
				listReturnMap[in[0]] = list[in[0]]
			}
		}
		for l := 0; l < length; l++ {
			listReturn = append(listReturn, listReturnMap[l])
		}
		return listReturn
	}
	return list
}

func NotConcurency(list []int64, max int) []int64 {
	var (
		i          int
		listReturn []int64

		j      = max
		length = len(list)
	)

	if length > max {
		for j <= length && i < length {
			time.Sleep(1 * time.Millisecond)
			listReturn = append(listReturn, list[i:j]...)
			i += max
			j += max
			if j > length {
				j = length
			}
		}
	} else {
		listReturn = list
	}
	return listReturn
}
