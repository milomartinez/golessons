package players

import (
	"fmt"
)

//Info :Datos del usuario
type Info struct {
	Name    string
	Age     int
	Balance int64
	casino  string
}

//New : Contructor
func New(sname string, nage int, nbalance int64, scasino string) Info {

	info := Info{sname, nage, nbalance, scasino}
	return info
}

//Print : print all values
func (i *Info) Print() {

	fmt.Println(i)

}

//GetCasino : get casino value
func (i *Info) GetCasino() string {

	return i.casino
}
