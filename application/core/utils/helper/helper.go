package helper

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

func Contains(s []any, e string) bool {
	for _, a := range s {
		if a.(string) == e {
			return true
		}
	}
	return false
}

func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ItemsWithPrefixInRole(s []any, prefix string) []string {
	orgs := []string{}
	for _, a := range s {
		if strings.HasPrefix(a.(string), prefix) {
			orgs = append(orgs, strings.ToLower(strings.TrimPrefix(a.(string), prefix)))
		}
	}
	return orgs
}

func Remove[T comparable](l []T, item T) []T {
	for i, elem := range l {
		if elem == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

func RemoveByIndex[T any](l []T, i int) []T {
	return append(l[:i], l[i+1:]...)
}

func Minimum(a uint, b uint) uint {
	if a < b {
		return a
	}
	return b
}

func HMDays(day1 string, day2 string) int {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	indx1 := 0
	indx2 := 0
	for i, d := range days {
		if d == day1 {
			indx1 = i
			break
		}
	}

	for i, d := range days {
		if d == day2 {
			indx2 = i
			break
		}
	}
	fmt.Println(indx2)
	if indx1 > indx2 {

		return 7 - indx1 + indx2
	} else {
		return indx1 - indx2
	}
}
