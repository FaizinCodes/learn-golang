package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// tes print code
	fmt.Println("hello world", "how", "are", "your", "day")

	// variable on go
	var firstName string = "faiz"
	var lastName string
	lastName = "in tech"

	fmt.Printf("hello %s %s!\n", firstName, lastName)

	// declarasi variable tanpa type data
	name := "Apa kabar?"

	fmt.Printf("%s", name)

	// deklarasi multiple variable
	var one, two, three string
	one, two, three = "satu", "dua", "tiga"

	// atau
	var empat, lima, enam = "empat", "lima", "enam"

	// atau
	tujuh, lapan, sembilan := "tujuh", "lapan", "sembilan"

	// atau
	nomor, bol, flot, str := 1, true, 2.2, "test"

	fmt.Println(one, two, three, empat, lima, enam)
	fmt.Printf("test apa %s %s %s %d %t %f %s", tujuh, lapan, sembilan, nomor, bol, flot, str)

	// variable underscore "_" digunakan untuk menampung variable yang tidak di pakai
	_ = "belajar golang"
	_ = "golang itu mudah"
	haha, _ := "jhon", "wick"
	fmt.Println("\n" + haha)

	// variable menggunakan new
	tempat := new(string)
	fmt.Println(*tempat)

	// penggunaan function print
	// perbedaan println() dan print yaitu println memuat garis baru, sedangkan print tidak
	println(name)
	print(haha)
	print(one + "\n")
	// output
	// Apa kabar?
	// jhonsatu

	// kondisi if biasa
	var point = 9

	if point == 10 { //tidak menggunakan parenthese ()
		fmt.Println("sepuluh")
	} else if point < 10 {
		fmt.Println("sedang")
	} else {
		fmt.Println("kurang")
	}

	// variable temporary pada perulangan
	// varible yang bisa digunakan pada blok kondisi dimana dia di rempatkan saja
	var nilai = 8849.0

	if percent := nilai / 100; percent >= 100 {
		fmt.Printf("%.1f%s percent!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good!\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad!\n", percent, "%")
	}

	// kondisi menggunakan switch - case
	var no = 1

	switch no {
	case 9:
		fmt.Println("baguss")
	case 5:
		fmt.Println("kurang")
	default:
		fmt.Println("coba lagi")
	}

	// case dengan banyak kondisi
	switch no {
	case 10, 9, 8, 7:
		fmt.Println("dekat")
	case 5:
		fmt.Println("jauh")
	}

	// case dengan gaya if-else
	switch {
	case no > 10:
		fmt.Println("besar")
	case no <= 10 && no > 5:
		fmt.Println("ideal")
	default:
		{
			fmt.Println("apa ini")
			fmt.Println("jelek")
		}
	}

	// fallthroug untuk memaksaa proses pengecekan diteruskan ke case selanjutnya
	switch {
	case no == 8:
		fmt.Println("perfect")
	case no < 2 && no > 0:
		fmt.Println("awesome")
		fallthrough // case dibawah juga akan di baca
	case no < 5:
		fmt.Println("belajr lagi ya")
	default:
		{
			fmt.Println("coba lebih giat")
			fmt.Println("pastibisa")
		}
	}

	// perulangan
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// atau
	var i = 0

	for i < 5 {
		fmt.Println("angka", i)
		i++
	}

	// atau
	var o = 0
	for {
		fmt.Println("Angka", o)
		o++
		if o < 5 {
			break
		}
	}

	// atau
	for a := 1; a <= 10; a++ {
		if a%2 == 1 {
			continue //dipaksa ke angak selanjutnya tanpa mengkesekusi code dibawah
		}

		if a > 8 {
			break //code berhenti
		}

		fmt.Println("Angka", a)
	}

	//atau
	for z := 0; z < 5; z++ {
		for j := z; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	// atau
cobeloop:
	for d := 0; d < 5; d++ {
		for n := 0; n < 5; n++ {
			if d == 3 {
				break cobeloop
			}
			fmt.Print("matriks [", d, "][", n, "]", "\n")
		}
	}

	// array
	var names [4]string
	names[0] = "luffy"
	names[3] = "mugiwara"
	names[1] = "mongky"
	names[2] = "D"
	// names[4] = "wkwk" //invalid aray index out of bounds

	fmt.Println(names)

	// inisialisasi array value
	var fruit = [4]string{"apple", "manggo", "banana", "melon"}

	fmt.Println("jumlah buah = \t\t", len(fruit))
	fmt.Println("isi semua element \t", fruit)

	// atau
	var food [4]string
	food = [4]string{"noodles", "friedchicken", "friedrice", "minas"}

	fmt.Println(food)

	// insialisasi array tanpa nilai awal
	var somthing = [...]int{2, 3, 3, 3}

	fmt.Println("length \t:", len(somthing))
	fmt.Println("data array \t:", somthing)

	// multidimension array
	var numbers1 = [2][3]int{[3]int{1, 2, 3}, [3]int{2, 3, 4}}
	// atau
	var numbers2 = [2][3]int{{2, 5, 6}, {5, 6, 7}}

	fmt.Println(numbers1)
	fmt.Println(numbers2)

	// perulangan dengan array menggunakan for - len()
	var places = [4]string{"jogja", "padang", "medan", "jakarta"}

	for r := 0; r < len(places); r++ {
		fmt.Printf(`liburan di %s sangat menyenangkan`+"\n", places[r])
	}

	// perulangan di array dengan menggunakan for - range
	for h, place := range places {
		fmt.Printf("elemen %d : %s\n", h, place)
	}

	// underscore best practise uses
	for _, place := range places {
		fmt.Printf("lokasi : %s\n", place)
	}
	// atau
	for h, _ := range places {
		fmt.Printf("nomor : %d\n", h)
	}

	// alokasi element array menggunakan keyword make
	var anything = make([]string, 3)
	anything[0] = "I"
	anything[1] = "love"
	anything[2] = "my"
	// anything[3] = "?"

	fmt.Println(anything)

	// slice
	var tryslice = []string{"coba1", "coba2", "coba3"}

	fmt.Println(tryslice)      //semua element
	fmt.Println(tryslice[0:2]) //semua element mulai dari index 0 sampai index sebelum 2
	fmt.Println(tryslice[:])   //semua element
	fmt.Println(tryslice[:1])  //semua element hingga sebelum index 1
	fmt.Println(tryslice[2:])  //semua element dimulai dari index 2

	// slice merupakan tipe data reference
	// jika ada slice baru yang terbentuk dari slice yang lama maka data elemen slice baru memiliki alamat memory yang sama
	// dengan kata lain jika terjadi perubahan pada slice baru maka akan mempengaruhi slice lama

	var myfavoritefruit = []string{"apple", "banana", "dragonfruit", "waermellon", "lemon"}

	var firstfruit = myfavoritefruit[0:2]
	var secondfruit = myfavoritefruit[1:4]

	var newFruit = firstfruit[:1]
	var oldfruit = secondfruit[0:1]

	fmt.Println(myfavoritefruit) //[apple banana dragonfruit waermellon lemon]
	fmt.Println(firstfruit)      //[apple banana]
	fmt.Println(secondfruit)     //[banana dragonfruit waermellon]
	fmt.Println(newFruit)
	fmt.Println(oldfruit) //[banana]

	oldfruit[0] = "nanas" // akan mengganti kata banana

	fmt.Println(myfavoritefruit) //[apple nanas dragonfruit waermellon lemon]
	fmt.Println(firstfruit)      //[apple nanas]
	fmt.Println(secondfruit)     //[nanas dragonfruit waermellon]
	fmt.Println(newFruit)
	fmt.Println(oldfruit) //[nanas]

	// penggunaan len()
	// pada slice len digunakna untuk menghitung jumlah elemen yang ada dalam slice

	fmt.Println(len(myfavoritefruit)) // hasilnaya 5

	// penggunaan cap()
	// pada slice cap digunakan untuk menghitung kapasaitas maksimal dalam elemen slice

	fmt.Println(cap(myfavoritefruit)) // hasilnya 5
	fmt.Println(cap(firstfruit))      // haislnya tetap 5 karena diambil dari sumber slice yang sama
	fmt.Println(cap(newFruit))        // haislnya tetap 5 karena diambil dari sumber slice yang sama

	// penggunaan append()
	// pada slice append digunakan untuk menambahkan element pada slice

	var afterfruit = append(myfavoritefruit, "orange")
	fmt.Println(len(myfavoritefruit))
	fmt.Println(cap(myfavoritefruit))

	fmt.Println(myfavoritefruit)
	fmt.Println(afterfruit)

	var insertnewfruit = append(newFruit, "leci")
	fmt.Println(len(newFruit))
	fmt.Println(cap(newFruit))

	fmt.Println(myfavoritefruit)
	fmt.Println(insertnewfruit)

	// penggunaan copy()
	// digunakan slice untuk mengopy parameter ke 2 kepada paramater 1

	dst := make([]string, 3)
	src := []string{"watermellon", "penaple", "apple", "orange"}
	n := copy(dst, src) // akan mencopy ke dst dengan value [watermellon penaple apple] hanya 3 karena copy() hanya mengcopy element sebanyak len(dst)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

	// pengkakses element slice dengan 3 index
	// 3 index digunakan untuk slicing elemen sekaligus dengan menentukan kapasitasnya

	var afruits = myfavoritefruit[0:3]   // kapasitas maksimal sama dengan myfavoritefruit yaitu 5
	var bfruits = myfavoritefruit[0:3:3] // kpasitas maksimal 3

	fmt.Println(myfavoritefruit)
	fmt.Println(len(myfavoritefruit))
	fmt.Println(cap(myfavoritefruit))

	fmt.Println(afruits)
	fmt.Println(len(afruits))
	fmt.Println(cap(afruits))

	fmt.Println(bfruits)
	fmt.Println(len(bfruits))
	fmt.Println(cap(bfruits))

	// map
	// nilai default maf adalah nil
	var mounth = map[string]int{}
	fmt.Println(mounth)

	mounth["january"] = 10
	mounth["february"] = 20

	fmt.Println(mounth)

	// inisialisasi item map
	var animal = map[string]int{
		"horse":   10,
		"elephan": 10,
	} // dengan value

	var stars = map[string]int{}       // tanpa value
	var another = make(map[string]int) //tanpa value menggunakan make()
	var others = *new(map[string]int)  //tanpa value menggunakan new()

	fmt.Println(animal)
	fmt.Println(stars)
	fmt.Println(another)
	fmt.Println(others)

	// iterasi item map menggunakan for - range
	for key, val := range animal {
		fmt.Println(key, " \t", val)
	}

	// menghapus item map dengan delete()
	delete(animal, "horse")
	fmt.Println(len(animal))
	fmt.Println(animal)

	// cek ketersediaan elemen dalam map
	var value, hewan = animal["elephan"] // variable pertama berisi nilai element dan variable kedua berisi boolean

	if hewan {
		fmt.Println(value)
	} else {
		fmt.Println("Item does'nt exist")
	}

	// kombinasi slice dan map
	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	var datas = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}

	for _, data := range datas {
		fmt.Println(data["name"] + data["address"])
	}

	showMessage("halo", myfavoritefruit)

	rand.Seed(time.Now().Unix()) //mendapatka angka random
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)

	var diameter float64 = 15
	var area, circumference = calculated(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

	var ore, ga = calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", ore)
	fmt.Printf("keliling lingkaran\t: %.2f \n", ga)

	// penggunaan function closure
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max) //variable %v pada printf berguna untuk menampilkan semua jenis type data

	// penggunana function closure IIFE
	// Immediately-Invoke Function Expression
	// function closure ini dieksekusi langsung sesaat setelah di deklarasinkan

	var angkaasli = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range angkaasli {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println("original number :", angkaasli)
	fmt.Println("filtered number :", newNumbers)

}

// learn function
func showMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int { //function untuk mengembalikan nilai
	var nilai = rand.Int()%(max-min+1) + min
	return nilai
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return // akan memberhentikan proses function
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

// function multiple return
func calculated(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2) // math.pow() digunakan untuk memangkatkan nilai param 1 pangkat param 2

	// hitung keliling
	var circumference = math.Pi * d //math.pi merepresentasikan Pi atau 22/7

	// kembalikan 2 nilai
	return area, circumference // akan mengembalikan nilai untuk param 1 dan param 2

}

func calculate(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return
}
