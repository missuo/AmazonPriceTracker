/*
 * @Author: Vincent Young
 * @Date: 2023-10-15 22:53:26
 * @LastEditors: Vincent Young
 * @LastEditTime: 2023-10-16 02:07:01
 * @FilePath: /AmazonPriceTracker/main.go
 * @Telegram: https://t.me/missuo
 *
 * Copyright © 2023 by Vincent, All Rights Reserved.
 */

package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func pricer(productLink string) map[string]interface{} {
	var newPrice string
	var productTitle string
	var usedPrice string
	var savingsPercentage string
	productDetail := make(map[string]interface{})

	c := colly.NewCollector(
		colly.AllowedDomains("www.amazon.com"),
	)

	var callbackTriggeredTitle bool
	var callbackTriggeredPrice bool
	var callbackTriggeredPercentage bool

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36")
	})

	c.OnHTML("#productTitle", func(e *colly.HTMLElement) {
		if !callbackTriggeredTitle {
			productTitle = e.Text
			productTitle = strings.TrimSpace(productTitle)
			callbackTriggeredTitle = true
		}
	})

	c.OnHTML("#corePriceDisplay_desktop_feature_div > div.a-section.a-spacing-none.aok-align-center > span.a-size-large.a-color-price.savingPriceOverride.aok-align-center.reinventPriceSavingsPercentageMargin.savingsPercentage", func(e *colly.HTMLElement) {
		if !callbackTriggeredPercentage {
			savingsPercentage = e.Text
			callbackTriggeredPercentage = true
		}
	})

	c.OnHTML("#corePriceDisplay_desktop_feature_div > div.a-section.a-spacing-none.aok-align-center > span.a-price.aok-align-center.reinventPricePriceToPayMargin.priceToPay > span:nth-child(2)", func(e *colly.HTMLElement) {
		if !callbackTriggeredPrice {
			priceWhole := e.ChildText("span.a-price-whole")
			priceFraction := e.ChildText("span.a-price-fraction")
			newPrice = priceWhole + priceFraction
			callbackTriggeredPrice = true
		}
	})

	c.OnHTML("#olpLinkWidget_feature_div > div.a-section.olp-link-widget > span > a > div > div > span.a-price > span.a-offscreen", func(e *colly.HTMLElement) {
		usedPrice = e.Text
		usedPrice = strings.ReplaceAll(usedPrice, "$", "")
	})

	c.Visit(productLink)
	productDetail["product_title"] = productTitle
	productDetail["product_link"] = productLink
	productDetail["new_price"] = newPrice
	productDetail["used_price"] = usedPrice
	productDetail["savings_percentage"] = savingsPercentage
	return productDetail
}

func getProductId(link string) string {
	re := regexp.MustCompile(`/dp/([A-Z0-9]+)`)
	match := re.FindStringSubmatch(link)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func main() {
	fmt.Println("Amazon Price Tracking is running...")
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "Welcome to Amazon Price Tracker. Made by Vincent.",
		})
	})
	r.GET("/price", func(c *gin.Context) {
		rLink := c.Query("link")
		rId := c.Query("id")
		productId := rId
		if productId == "" {
			productId = getProductId(rLink)
		}
		if productId != "" {
			productLink := "https://www.amazon.com/dp/" + productId
			product := pricer(productLink)
			// fmt.Println(product["product_title"])
			if product["product_title"] == "" {
				c.JSON(http.StatusTooManyRequests, gin.H{
					"code":    http.StatusTooManyRequests,
					"message": "Too Many Requests",
				})
			} else {
				c.JSON(http.StatusOK, product)
			}
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "Link Not Found",
			})
		}

	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Path not found",
		})
	})

	r.Run(":7777")
}
