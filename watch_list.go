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

	for i := len(w.movies)-1; i>=0; i-- {
		if i == 0 {
			fmt.Printf("%s.\n\n", w.movies[i])
		} else {
			fmt.Printf("%s\n", w.movies[i])
		}
	}

}

func main() {
	Fav := NewList(3)
	Fav.Add("Avengers")
	Fav.Add("Titanic")
	Fav.Add("Star Trek")
	Fav.Print()
	// 	Star Trek
	// 	Titanic
	// 	Avengers.

	Fav.Add("Gundala")
	Fav.Print()
	// Gundala
	// Star Trek
	// Titanic.

	Fav.Watch("Titanic")
	Fav.Print()
	// Titanic
	// Gundala
	// Star Trek.

	Fav.Add("Justice League")
	Fav.Print()
	// Justice League
	// Titanic
	// Gundala.	
}
