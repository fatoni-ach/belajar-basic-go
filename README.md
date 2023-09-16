# Belajar Go Lang
Repository ini dibuat untuk catatan selama belajar bahasa go lang.
Untuk mencobanya kamu bisa menjalankan satu persatu file yang ada di root folder.<br>
Example : <br>
```shell
go run [nama file]
```
<br>

---
## **Basic**
Untuk menjalankan program bisa langkah pertama build script dan jalankan kan hasil build. berikut scriptnya :
```
go build [nama_file]
./[nama_file_hasil_build]

ex : 
go build konversi.go
./konversi
```
dalam proses developmet bisa menggunakan perintah `run` namun tidak disarankan digunakan di production environment karena pada akhirnya yang di kirim ke server adalah hasil build an. berikut scriptnya :
```
go run [nama_file]

ex : 
go run konversi.php
```

### **Array dan Slice**
Array dan slice hampir sama, yang membedakan slice lebih dynamic dibandingkan array dan kedepannya lebih sering menggunakan type data Slice. perbedaan deklarasi array dan slice adalah seperti berikut :
```go
iniArray := [...]int{1,2,3,4,5}

iniSLice := []int{1,2,3,4,5}
```

### **Named Return values**
Named return value hampir sama seperti fungsi biasa, yang membedakan dengan ini bisa meng deklare return variable di sebelah function parameter dan hanya perlu meng return tanpa harus menuliskan variable yang di return secara explisit (namun bisa juga dituliskan ulang). contohnya seperti berikut :
```go
func getFullNameWithNamedValue() (firstname string, lastname string, age int) {
	firstname = "Achmad"
	lastname = "fatoni"
    age = 10
    // variable yang di return tidak perlu ditulis kembali 
	return
}
``` 

### **Variadic Function**
Merupakan function yang terdapat parameter yang menjadi array. Parameter tersebut wajib ditaruh di paling akhir diikuti `...` (titik 3) di sebelum tipe data. Parameter tersebut disebut sebagai varargs yang berarti datanya bisa menerima lebih dari 1 input yang akan dianggap sebagai array / slice. yang membedakan dengan array biasa kita wajib membuat array terlebih dahulu sebelum mengirimkannya ke function sedangkan dengan varargs kita bisa langsung mengirimkan datanya. Conntohnya seperti berikut. 
```go
func sumAll(numbers ...int) (result int) {
	result = 0
	for _, number := range numbers {
		result = result + number
	}
	return
}

func main() {
	total := sumAll(10, 10, 10)
	fmt.Println("Total : ", total)

	// mengirimkan data yang sudah menjadi array ke varargs
	numbers := []int{10, 10, 10, 10}
	fmt.Println("Total : ", sumAll(numbers...))
}
```
### **Struct Function**
Struct function merupakan tipe data yang mirip seperti tipe data lainnya. Namun kita bisa menambahkan function ke dalam struct yang seakan-akan function tersebut punya si struct yang sebenarnya function tersebut bukan punya si struct atau berdiri sendiri.
Contohnya bisa  dilihat di file berikut : 
[Click Here](/struct-function.go)
<br>
### **Interface**
Merupakan tipe data abstrak, dia tidak memiliki implementasi langsung. Interface berisikan definisi method dan biasa digunakan sebagai kontrak. Contohnya bisa dilihat di file berikut : [Click Here](/interface.go)
<br>

### **Package *strings***
Package ini digunakan untuk memanipulasi string. Contohnya seperti berikut :
```go
func main() {
	word := "Hallo semuanya"

	fmt.Println("word : ", word)
	fmt.Println("word UpperCase : ", strings.ToUpper(word))
	fmt.Println("word LowerCase : ", strings.ToLower(word))
}
```
### **Package sort**
Package ini digunakan untuk proses sorting. Contohnya bisa diihat [disini](sort.go)