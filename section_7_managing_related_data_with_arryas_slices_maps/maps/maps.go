package main

import "fmt"

//difference strcts and maps is that maps store key-value pairs

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.amazon.com",
	}
	fmt.Println(websites)

	//advantage you have constant time complexity O(1) to access any value by its key
	fmt.Println(websites["Amazon Web Services"])

	//with maps you can add new key-value pairs at any time
	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	//overwrite existing value
	websites["Google"] = "https://google.co.uk"
	fmt.Println(websites)

	//delete key-value pair
	delete(websites, "Linkedin")
	fmt.Println(websites)

}
