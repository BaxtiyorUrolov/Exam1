package task

import (
	"fmt"
	"math"
	"modul/read"
	"sort"
)

func Task1(info []read.Info) {
	cashSum := 0
	totalSum := 0

	for _, item := range info {
		cashSum += item.Cash
		totalSum += item.Basket.Total
	}

	fmt.Println("Umumiy mablag': ", cashSum)
	fmt.Println("Umumiy harajat: ", totalSum)
}

func Task2(info []read.Info) {
	sort.Slice(info, func(i, j int) bool {
		return info[i].Basket.Total > info[j].Basket.Total
	})

	for i, customer := range info {
		if i >= 1 {
			break
		}
		fmt.Printf("Eng ko'p pul sarf etgan mijoz: %s %s(%dso'm)\n", customer.FirstName, customer.LastName, customer.Basket.Total)
	}
}

func Task3(info []read.Info) {
	var allProducts []read.Product

	for _, customer := range info {
		allProducts = append(allProducts, customer.Basket.Product...)
	}

	sort.Slice(allProducts, func(i, j int) bool {
		return allProducts[i].Price > allProducts[j].Price
	})

	if len(allProducts) > 0 {
		maxProduct := allProducts[0]
		fmt.Printf("Eng qimmat mahsulot: %s, Narxi: %d\n", maxProduct.Name, maxProduct.Price)
	}
}

func Task4(info []read.Info) {
	var allProducts []read.Product

	for _, customer := range info {
		for _, product := range customer.Basket.Product {
			allProducts = append(allProducts, product)
		}
	}

	sum := 0
	for _, product := range allProducts {
		sum += product.Price
	}

	averagePrice := 0
	if len(allProducts) > 0 {
		averagePrice = sum / len(allProducts)
	}

	fmt.Println("Barcha mahsulotlarning urtacha narxi: ", averagePrice)
}

func Task5(info []read.Info) {
	sort.Slice(info, func(i, j int) bool {
		return info[i].Basket.Total < info[j].Basket.Total
	})

	minSpender := info[0]

	fmt.Printf("Eng arzon savdo qilgan mijoz: %s %s, Jami xarajat: %d\n", minSpender.FirstName, minSpender.LastName, minSpender.Basket.Total)
}

func Task6(info []read.Info) {
	categoryTotal := make(map[string]int)

	for _, customer := range info {
		for _, product := range customer.Basket.Product {
			categoryTotal[product.Category] += product.Quantity
		}
	}

	maxCategory := ""
	maxQuantity := 0
	for category, quantity := range categoryTotal {
		if quantity > maxQuantity {
			maxQuantity = quantity
			maxCategory = category
		}
	}

	if maxQuantity > 0 {
		fmt.Printf("Eng ko'p sotilgan kategoriya: %s, Sotilgan miqdori: %d\n", maxCategory, maxQuantity)
	}
}

func Task7(info []read.Info) {
	var maxSpender read.Info

	var allProducts []read.Product

	for _, customer := range info {
		for _, product := range customer.Basket.Product {
			allProducts = append(allProducts, product)
		}
	}

	sort.Slice(info, func(i, j int) bool {
		return allProducts[i].Quantity > allProducts[j].Quantity
	})

	if len(info) > 0 {
		maxSpender = info[0]

		// Natijani chiqaramiz
		fmt.Printf("Eng ko'p sotilgan mahsulot: %s, soni: %d\n", maxSpender.Basket.Product[0].Name, maxSpender.Basket.Product[0].Quantity)
	}
}

func Task8(info []read.Info) {
	var allProducts []read.Product

	for _, customer := range info {
		for _, product := range customer.Basket.Product {
			allProducts = append(allProducts, product)
		}
	}

	sum := 0
	for _, product := range allProducts {
		sum += product.Quantity
	}

	averageQuantity := 0
	if len(allProducts) > 0 {
		averageQuantity = sum / len(allProducts)
	}

	fmt.Println("Barcha mahsulotlarning o'rtacha miqdori: ", averageQuantity)
}

func Task9(info []read.Info) {
	var maxSpender read.Info

	sort.Slice(info, func(i, j int) bool {
		return len(info[i].Basket.Product) > len(info[j].Basket.Product)
	})

	if len(info) > 0 {
		maxSpender = info[0]

		fmt.Printf("Eng ko'p mahsulot sotib olgan mijoz: %s %s (sotib olgan mahsulotlar soni: %d)\n", maxSpender.FirstName, maxSpender.LastName, len(maxSpender.Basket.Product))
	}
}

func Task10(info []read.Info) {
	var (
		firstname10 string
		lastname10  string
		productname string
		max         int = math.MinInt
	)
	for _, value := range info {
		for _, value2 := range value.Basket.Product {
			if max < value2.Quantity {
				max = value2.Quantity
				firstname10 = value.FirstName
				lastname10 = value.LastName
				productname = value2.Name
			}
		}
	}

	fmt.Println("Eng ko'p savdolarda ko'ringan mahsulot: ",productname, "(", firstname10, lastname10, "-", max, ")")
}

func Task11(info []read.Info) {
	totalSum := 0

	for _, item := range info {
		totalSum += item.Basket.Total
	}

	sort.Slice(info, func(i, j int) bool {
		return info[i].Cash > info[j].Cash
	})

	fmt.Println("O'rtacha savdo mablag'i: ", totalSum/len(info))

	for i, customer := range info {
		if i >= 1 {
			break
		}
		fmt.Printf("Eng ko'p mablag'ga ega mijoz: %s %s (%dso'm)\n", customer.FirstName, customer.LastName, customer.Cash)
	}
}

func Task12(info []read.Info) {
	var maxTotal int
	var bestProduct read.Product

	for _, customer := range info {
		for _, product := range customer.Basket.Product {
			total := product.Quantity * product.Price
			if total > maxTotal {
				maxTotal = total
				bestProduct = product
			}
		}
	}

	if maxTotal > 0 {
		fmt.Printf("Eng ko'p foydalanilgan mahsulot: %s, kategoriya: %s, umumiy qiymati: %d\n", bestProduct.Name, bestProduct.Category, maxTotal)
	} else {
		fmt.Println("Mahsulot topilmadi")
	}
}

func Task13(info []read.Info) {
	for _, user := range info {
		max := math.MinInt
		var found bool
		var maxProduct read.Product
		for _, product := range user.Basket.Product {
			if price := product.Price * product.Quantity; price > max {
				max = price
				found = true
				maxProduct = product
			}
		}
		if found {
			fmt.Println(maxProduct.Name, " (",user.FirstName, user.LastName," - ", max, ")")
		}
	}
}

func Task14(info []read.Info) {
	for _, customer := range info {
		categoryTotal := make(map[string]int)

		for _, product := range customer.Basket.Product {
			categoryTotal[product.Category] += product.Quantity
		}

		maxCategory := ""
		maxQuantity := 0
		for category, quantity := range categoryTotal {
			if quantity > maxQuantity {
				maxQuantity = quantity
				maxCategory = category
			}
		}

		if maxQuantity > 0 {
			fmt.Printf("%s %s uchun eng ko'p sotilgan kategoriya: %s, sotilgan miqdori: %d\n",
				customer.FirstName, customer.LastName, maxCategory, maxQuantity)
		}
	}
}

func Task15(info []read.Info) {
	totalQuantity := 0
	for _, item := range info {
		for _, product := range item.Basket.Product {
			totalQuantity += product.Quantity
		}
	}

	fmt.Println("Umumiy mahsulotlar soni: ", totalQuantity)
}
