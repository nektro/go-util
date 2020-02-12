package stringsu

import (
	mapset "github.com/deckarep/golang-set"
)

func Depupe(a []string) []string {
	return fixArray2(mapset.NewSet(fixArray1(a)...).ToSlice())
}

func fixArray1(a []string) []interface{} {
	r := []interface{}{}
	for _, item := range a {
		r = append(r, item)
	}
	return r
}

func fixArray2(a []interface{}) []string {
	r := []string{}
	for _, item := range a {
		r = append(r, item.(string))
	}
	return r
}
