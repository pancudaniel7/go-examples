package internal

import "strings"

func SimpleString(value string) string {
	return value
}

func UpperCase(value string) string {
	return strings.ToUpper(value)
}

func LowerCase(value string) string {
	return strings.ToLower(value)
}

func Concat(value string, conct string) string {
	var sb strings.Builder
	sb.WriteString(value)
	sb.WriteString(conct)
	return sb.String()
}

func Contains(value string, sub string) bool {
	return strings.Contains(value, sub)
}

func GetByIndex(value string, index int) any {
	return value[index]
}

func Replace(value string, old string, new string) string {
	return strings.Replace(value, old, new, -1)
}

func Slicing(value string, start int, end int) string {
	return value[start:end]
}

func HasPrefix(value string, prefix string) bool {
	return strings.HasPrefix(value, prefix)
}

func HasSuffix(value string, suffix string) bool {
	return strings.HasSuffix(value, suffix)
}

func Split(value string, sep string) []string {
	return strings.Split(value, sep)
}

func Join(elements []string, sep string) string {
	return strings.Join(elements, sep)
}

func Index(value string, sub string) int {
	return strings.Index(value, sub)
}

func LastIndex(value string, sub string) int {
	return strings.LastIndex(value, sub)
}

func Repeat(value string, count int) string {
	return strings.Repeat(value, count)
}

func Trim(value string, cutset string) string {
	return strings.Trim(value, cutset)
}

func TrimSpace(value string) string {
	return strings.TrimSpace(value)
}

func Fields(value string) []string {
	return strings.Fields(value)
}

func Compare(a, b string) int {
	return strings.Compare(a, b)
}

func EqualFold(a, b string) bool {
	return strings.EqualFold(a, b)
}

func Count(value string, sub string) int {
	return strings.Count(value, sub)
}
