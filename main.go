package main

import (
	"context"
	"fmt"
	"os"

	"github.com/web3-storage/go-w3s-client"
)

func main() {

	// Tokenimizi ekliyoruz.
	conn, err := w3s.NewClient(w3s.WithToken("TOKEN"))

	if err != nil {
		panic(err)
	}

	//Boş bir durum döndürüyoruz.
	ctx := context.Background()

	// Seçtiğimiz resim mi açıyoruz.
	img, _ := os.Open("./ornek.png")

	// Tokenimize durumu atıyoruz ve içini resimimizi ekliyoruz. sonuç olarak cid yani benzersiz bir parmak izi döndürüyor.
	cid, _ := conn.Put(ctx, img)

	// Url erişim çıktımızı alıyoruz ve tarayıcıya yapıştırarak erişim sağlıyoruz.
	fmt.Printf("https://ipfs.io/ipfs/%s\n", cid)
}
