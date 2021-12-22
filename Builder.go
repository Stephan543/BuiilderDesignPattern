package main

import (
	"fmt"
	"math/rand"
)

// import "fmt"

type iBuilder interface {
	setAdType()
	setAdCountry()
	setAdId(adId int)
	getAd() ad
}

type ad struct {
	adType    string
	adCountry string
	adId      int
}

func getBuilder(builderType string) iBuilder {
	if builderType == "img" {
		return &imgAdBuilder{}
	}

	if builderType == "video" {
		return &videoAdBuilder{}
	}

	return nil

}

/*			Image Add Builder				*/

type imgAdBuilder struct {
	adType    string
	adCountry string
	adId      int
}

//func newImgAdBuilder() *imgAdBuilder {
//	return &imgAdBuilder{}
//}

func (b *imgAdBuilder) setAdType() {
	b.adType = "Image Ad"
}

func (b *imgAdBuilder) setAdCountry() {
	b.adCountry = "Canada"
}

func (b *imgAdBuilder) setAdId(adId int) {
	b.adId = adId
}

func (b *imgAdBuilder) getAd() ad {
	return ad{
		adType:    b.adType,
		adCountry: b.adCountry,
		adId:      b.adId,
	}
}

/*			Image Add Builder				*/

type videoAdBuilder struct {
	adType    string
	adCountry string
	adId      int
}

//func newVideoAdBuilder() *videoAdBuilder{
//	return &videoAdBuilder{}
//}

func (b *videoAdBuilder) setAdType() {
	b.adType = "Video Ad"
}

func (b *videoAdBuilder) setAdCountry() {
	b.adCountry = "USA"
}

func (b *videoAdBuilder) setAdId(adId int) {
	b.adId = adId
}

func (b *videoAdBuilder) getAd() ad {
	return ad{
		adType:    b.adType,
		adCountry: b.adCountry,
		adId:      b.adId,
	}
}

/*			Director Builder				*/

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildAd(adId int) ad {
	d.builder.setAdType()
	d.builder.setAdCountry()
	d.builder.setAdId(adId)
	return d.builder.getAd()
}

func main() {

	imgAdBuilder := getBuilder("img")
	videoAdBuilder := getBuilder("video")

	/*			Img Ad Init				*/
	director := newDirector(imgAdBuilder)
	imgAd := director.buildAd(rand.Intn(100000000))

	fmt.Printf("Ad Type: %s\n", imgAd.adType)
	fmt.Printf("Ad Country %s\n", imgAd.adCountry)
	fmt.Printf("Ad Id %d\n", imgAd.adId)

	/*			Video Ad Init				*/

	director.setBuilder(videoAdBuilder)
	videoAd := director.buildAd(rand.Intn(1000000000))

	fmt.Printf("\nAd Type: %s\n", videoAd.adType)
	fmt.Printf("Ad Country %s\n", videoAd.adCountry)
	fmt.Printf("Ad Id %d\n", videoAd.adId)
}
