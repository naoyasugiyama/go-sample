package somepkg

/*
	大文字はパブリックメンバー他のファイルからも import すればアクセス可能
*/

var SomeVar int
var someVar2 int

func SomeFunc() {
	SomeVar = 10
	someVer2 = 1
}

func someFunc2() {
	SomeFunc()
}

func somepkg() {

}

func Swap(x, y string) (string, string) {
	return y, x
}
