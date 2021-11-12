package main

import "fmt"

type WatchList struct {
	capacity int
	movies   []string
}

func NewList(cap int) WatchList {
	list := make([]string, cap)
	return WatchList{
		capacity: 0,
		movies:   list,
	}
}

func (w *WatchList) findMovieIndex(movie string) int {
	for i, v := range w.movies {
		if v == movie {
			return i
		}
	}
	return -1
}

func (w *WatchList) shift() {
	for i := 0; i < w.capacity-1; i++ {
		w.movies[i] = w.movies[i+1]
	}
}

func (w *WatchList) Add(movie string) {
	moviesSize := len(w.movies)
	if moviesSize == w.capacity {
		w.shift()
		w.movies[w.capacity-1] = movie
	} else {
		w.movies[w.capacity] = movie
		w.capacity = w.capacity + 1
	}
}

func (w *WatchList) Watch(movie string) {
	w.Add(movie)
}

func (w *WatchList) Print() {

	for i, v := range w.movies {
		if i < len(w.movies)-1 {
			fmt.Printf("%s, ", v)
		} else {
			fmt.Printf("%s\n", v)
		}
	}

}

func main() {
	Fav := NewList(3)
	Fav.Add("Avengers")
	Fav.Add("Titanic")
	Fav.Add("Star Trek")
	Fav.Print()
	//  Avengers, Titanic, Star Trek

	Fav.Add("Gundala")
	Fav.Print()
	// Titanic, Star Trek, Gundala

	Fav.Watch("Titanic")
	Fav.Add("Justice League")
	Fav.Print()
	// Gundala, Titanic, Justice League
}
