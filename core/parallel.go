package core

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/hashicorp/go-retryablehttp"
)

type Info struct {
	url          string
	location     string
	length       int64
	pieces       int
	piece_length int64
}

type piece struct {
	index int
	begin int64
	end   int64
	done  bool
}

func Make_Info(url string, location string, pieces int) Info {
	len := get_length(url)
	return Info{
		url:          url,
		location:     location,
		length:       len,
		pieces:       pieces,
		piece_length: len / int64(pieces),
	}
}

func get_length(url string) int64 {
	res, err := http.Head(url)
	if err != nil {
		panic(err)
	}
	return res.ContentLength
}

func make_pieces(inf Info) []piece {
	pieces := make([]piece, inf.pieces)
	for i := 0; i < inf.pieces; i++ {
		if i == 0 {
			pieces[i].begin = 0
		} else {
			pieces[i].begin = pieces[i-1].end + 1
		}
		pieces[i].index = i
		pieces[i].done = false
		if i == inf.pieces-1 {
			pieces[i].end = inf.length - 1
		} else {
			pieces[i].end = pieces[i].begin + inf.piece_length - 1
		}
	}
	return pieces
}

func donwload_piece(wg *sync.WaitGroup, p *piece, inf Info) {
	defer wg.Done()
	// create file in tmp inside current directory
	// get current directory
	dir, _ := os.Getwd()
	temp_dir := dir + "/tmp"
	// check if tmp directory exists
	if _, err := os.Stat(temp_dir); os.IsNotExist(err) {
		// create tmp directory
		os.Mkdir(temp_dir, 0755)
	}
	addr := temp_dir + "/dat" + fmt.Sprint(p.index)
	f, err := os.Create(addr)
	if err != nil {
		fmt.Printf("Error creating file %s\n", addr)
		panic(err)
	}
	defer f.Close()
	// print path of file
	fmt.Println(f.Name())

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10
	retryClient.Logger = nil
	client := retryClient.StandardClient()

	req, _ := http.NewRequest("GET", inf.url, nil)
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", p.begin, p.end))
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	size, _ := io.Copy(f, res.Body)
	defer f.Close()
	p.done = true
	fmt.Printf("Piece %s downloaded, size: %d\n", addr, size)

}

func merge(inf Info, pieces []piece) error {
	f, err := os.Create(inf.location)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()
	for i := 0; i < inf.pieces; i++ {
		addr := "./tmp/dat" + fmt.Sprint(i)
		p, err := os.Open(addr)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer p.Close()
		io.Copy(f, p)
	}
	return nil
}

func Download(inf Info) error {
	pieces := make_pieces(inf)
	var wg sync.WaitGroup
	for i := 0; i < inf.pieces; i++ {
		wg.Add(1)
		go donwload_piece(&wg, &pieces[i], inf)
	}
	// wait for all pieces to be downloaded
	wg.Wait()

	fmt.Println("Downloaded all pieces")
	err := merge(inf, pieces)
	os.RemoveAll("./tmp")
	if err != nil {
		return err
	}
	fmt.Println("Downloaded")
	return nil
}
