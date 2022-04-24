package main

import (
	"flag"
	"fmt"
	"os"
  "math"
  "./kdoc"
)

var VERSION="0-3.1"
type Kontur struct {
	n         float64
	ro        float64
	r_kont    float64
	nu        float64
	l         float64
	r_zazeml  float64
	d         float64
	k_kratn   float64
}
func (k *Kontur) rz() float64 {
    ln:=math.Log(4*k.l/k.d)
//    fmt.Println(ln)
    k.r_zazeml=(k.ro/(6.28*k.l))*ln
    fmt.Println(k.r_zazeml," сопротивление 1-го заземлителя")
    return k.r_zazeml
}
func (k *Kontur) nf(){
    k.n=k.r_zazeml/(k.nu*k.r_kont)
    fmt.Println(int(k.n)+1,"заземлителей")
}
func (k *Kontur) area() float64 {
    return k.l*k.l*k.n
}
func (k *Kontur) len() float64 {
    return k.l*k.n
}

var  k=new(Kontur)
func main() {
  kl  := flag.Float64("l", 1.5, "длина заземлителя,м")
  knu := flag.Float64("nu", 0.42, "коэф. экранирования  заземлителя")
  kkr := flag.Float64("kr", 1, "коэф. кратности длины заземлителя к расстоянию между заземлителями (шага)")
  kro := flag.Float64("ro", 300, "удельное сопротивление грунта, ом/м")
  kd  := flag.Float64("d", 0.063, "диаметр заземлителя,м")
  kr_kont:= flag.Float64("R", 4, "сопротивление контура,ом")
  k_doc := flag.String("doc", "", "grunt грунт")
	flag.Parse()
  if (*k_doc=="grunt"||*k_doc=="грунт"){
      fmt.Println(kdoc.Grunt())
      os.Exit(0)
  }
	k.l     =* kl
	k.nu    =* knu
	k.k_kratn =* kkr
	k.ro    =* kro
	k.d     =* kd
	k.r_kont=* kr_kont
  fmt.Println("version",VERSION)
  fmt.Println(flag.Args())
	fmt.Println(k.l, "м длина заземлителя")
	fmt.Println(k.nu, "коэф. экранирования заземлителя")
	fmt.Println(k.k_kratn, "коэф. кратности шага заземлителей")
	fmt.Println(k.ro, "ом/м удельное сопротивление грунта")
	fmt.Println(k.d, "м диаметр заземлителя")
  fmt.Println(k.r_kont, "ом сопротивление контура")
  k.rz()
  k.nf()
  fmt.Println(int(k.len())+1,"м общая длина заземлителей")
  fmt.Println(int(k.area())+1,"м2 занимаемая площадь")
}
