package iteration

const num = 5
func Repeat(a string) string{
	var repeated string
	for i:=0; i < num; i++ {
		repeated += a
	}
	return repeated
}