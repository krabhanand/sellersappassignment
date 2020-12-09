package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"unicode"

	"sellerapps/internal/models"

	"github.com/PuerkitoBio/goquery"
)

func scrape(URL string) (models.Product, error) {
	product, err := GetLatestBlogTitles(URL)
	if err != nil {
		return models.Product{}, err
	}
	//	fmt.Printf(blogTitles)
	return product, err
}

// GetLatestBlogTitles gets the latest blog title headings from the url
// given and returns them as a list.
func GetLatestBlogTitles(url string) (models.Product, error) {

	// Get the HTML
	resp, err := http.Get(url)
	if err != nil {
		return models.Product{}, err
	}

	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return models.Product{}, err
	}

	title := ""
	doc.Find("title").Each(func(i int, s *goquery.Selection) {
		title += s.Text()
	})
	if strings.Contains(title, "Not Found") || !strings.Contains(title, "Amazon.") {
		return models.Product{}, errors.New(" this is not an amazon products page")
	}
	// Save each .post-title as a list
	//	titles := ""

	productTitle := ""
	doc.Find("#productTitle").Each(func(i int, s *goquery.Selection) {
		productTitle += s.Text()
	})

	productReviews := ""
	doc.Find("#acrCustomerReviewText").Each(func(i int, s *goquery.Selection) {
		productReviews += s.Text()
	})

	productPrice := ""
	doc.Find("#olp-sl-new-used").Each(func(i int, s *goquery.Selection) {
		//fmt.Println("inside finding product price")
		productPrice += s.Text()
	})
	productPrice = SpaceStringsBuilder(productPrice)
	productPrice = getPriceFromList(productPrice)

	productDescription := ""
	doc.Find("#productDescription").Each(func(i int, s *goquery.Selection) {
		productDescription += s.Text() + "\n"
	})
	productDescription = SpaceStringsBuilder(productDescription)

	landingImage := ""
	doc.Find("#imgTagWrapperId").Each(func(i int, s *goquery.Selection) {
		landingImage = s.Find("img").AttrOr("data-old-hires", "")
	})
	//landingImage = SpaceStringsBuilder(landingImage)

	productTitle = SpaceStringsBuilder(productTitle)
	productReviews = getReviewsNumber(productReviews)
	var product *models.Product = new(models.Product)
	product.Name = productTitle
	product.Description = productDescription
	product.Price = productPrice
	product.Reviews = productReviews
	product.ImageURL = landingImage
	fmt.Println("Product name: " + productTitle)
	fmt.Println("Product Reviews " + productReviews)
	fmt.Println("Product price: " + productPrice)
	fmt.Println("Product description: " + productDescription)
	fmt.Println("Image URL: " + landingImage)

	//titles += productTitle
	return *product, nil
}
func SpaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	prevspace := true
	//function to remove extra spaces
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
			prevspace = false
		} else if !prevspace {
			b.WriteRune(ch)
			prevspace = true
		}
	}
	var s string = b.String()
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}
func getReviewsNumber(str string) string {
	testArray := strings.Fields(str)
	return testArray[0]
}

func getPriceFromList(str string) string {
	testArray := strings.Fields(str)
	if len(testArray) > 0 {
		return testArray[len(testArray)-1]
	} else {
		return ""
	}
}
