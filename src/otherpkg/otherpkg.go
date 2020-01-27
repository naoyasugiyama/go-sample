package otherpkg

import "fmt"

/*
packageとimport
GOROOTがgoインストール直下になっている
otherpkg.go:6:8: cannot find package "somepkg" in any of:
        c:\go\src\somepkg (from $GOROOT)
		C:\Users\gotof\go\src\somepkg (from $GOPATH)

package 

$GOPATH/
	src/ 
		somepkg/
			some.go // package somepkg
			some_test.go // package somepkg_test

			// _test を入れるとテストケースとして扱われる
:禁止事項
	同じフォルダ内に複数のパッケージ宣言はできない
	循環参照できない
 
*/


//import "somepkg" // 別ファイル somepkgをインポート
import some "somepkg/somepkg" //更に別名宣言
// import . "somepkg" // メンバへのアクセス時にパッケージ名の指定を省略することが可能
/*
// error
func OtherFunc() {
	SomeFunc()
	SomeVer = 5
}*/

func OtherFunc() {
	some.SomeFunc()
	fmt.Println("somepkg.SomeVer:", somepkg.SomeVer)
	some.SomeVer = 10
}

