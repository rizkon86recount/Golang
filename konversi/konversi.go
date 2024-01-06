/*
// Atoi
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "42"
	integer, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Terjadi kesalahan:", err)
		return
	}
	fmt.Printf("Hasil konversi: %d\n", integer)
} */

/*
// ParseFloat
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "3.14"
	flt, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Terjjadi kesalahan:", err)
		return
	}
	fmt.Printf("Hasil konversi: %f\n", flt)
} */

/*
// Itoa
package main

import (
	"fmt"
	"strconv"
)

func main() {
	integer := 42
	str := strconv.Itoa(integer)
	fmt.Printf("Hasil Konversi: %s\n", str)
} */

/*
// FormatFloat
package main

import (
	"fmt"
	"strconv"
)

func main() {
	flt := 3.14159265359
	str := strconv.FormatFloat(flt, 'f', 2, 64)
	fmt.Printf("Hasil konversi: %s\n", str)
} */

/*
// strconv.ParseInt()
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "42"
	integer, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("Terjadi kesalahan:", err)
		return
	}
	fmt.Printf("Hasil konversi: %d\n", integer)
} */

/*
// strconv.FormatInt()
package main

import (
	"fmt"
	"strconv"
)

func main() {
	integer := int64(42)
	str := strconv.FormatInt(integer, 10)
	fmt.Printf("Hasil konversi: %s\n", str)
} */

/*
// Contoh penggunaan strconv.ParseFloat()
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "3.14"
	flt, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Terjadi kesalahan:", err)
		return
	}
	fmt.Printf("Hasil konversi: %f\n", flt)
} */

/*
// Contoh penggunaan strconv.FormatFloat():
package main

import (
	"fmt"
	"strconv"
)

func main() {
	flt := 3.14159265359
	str := strconv.FormatFloat(flt, 'f', 2, 64)
	fmt.Printf("Hasil konversi: %s\n", str)
} */

/*
// Contoh penggunaan strconv.ParseBool():
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1 := "true"
	boolValue1, err1 := strconv.ParseBool(str1)
	if err1 != nil {
		fmt.Println("Terjadi kesalahan:", err1)
	} else {
		fmt.Printf("Hasil konversi: %v\n", boolValue1)
	}

	str2 := "false"
	boolValue2, err2 := strconv.ParseBool(str2)
	if err2 != nil {
		fmt.Println("Terjadi kesalahan:", err2)
	} else {
		fmt.Printf("Hasil konversi: %v\n", boolValue2)
	}
} */

// Contoh penggunaan strconv.FormatBool():
package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolValue1 := true
	str1 := strconv.FormatBool(boolValue1)
	fmt.Printf("Hasil konversi: %s\n", str1)

	boolValue2 := false
	str2 := strconv.FormatBool(boolValue2)
	fmt.Printf("Hasil konversi: %s\n", str2)
}
