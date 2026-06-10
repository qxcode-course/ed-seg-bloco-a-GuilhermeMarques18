package main

func main() {
	queue := NewQueue[string]()

	for i := 'A'; i <= 'P'; i++ {
		queue.Enqueue(string(i))
	}

	queue.Torneio()
}
